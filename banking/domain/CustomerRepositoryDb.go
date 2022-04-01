package domain

import (
	"banking/errs"
	"banking/logger"
	"banking/sqlLogAdapter"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/sirupsen/logrus"
	"time"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, error) {
	//var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)
		//rows, err = d.client.Query(findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
		//rows, err = d.client.Query(findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, err
	}

	//err = sqlx.StructScan(rows, &customers)
	//if err != nil {
	//	logger.Error("Error while scanning customers " + err.Error())
	//	return nil, err
	//}

	//for rows.Next() {
	//	var c Customer
	//	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	//	if err != nil {
	//		logger.Error("Error while scanning customers " + err.Error())
	//		return nil, err
	//	}
	//	customers = append(customers, c)
	//}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	var c Customer
	err := d.client.Get(&c, customerSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &c, nil
}

func NewCustomeRepositoryDb() CustomerRepositoryDb {

	dsn := "root:1234@tcp(localhost:3306)/banking"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	logger := logrus.New()
	logger.Level = logrus.InfoLevel
	logger.Formatter = &logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	}
	loggerAdapter := sqlLogAdapter.NewLogrusAdapter(logger)

	//zapCfg := zap.NewProductionConfig()
	//zapCfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel) // whatever minimum level
	//zapCfg.DisableCaller = true
	//logger, _ := zapCfg.Build()
	//loggerAdapter := sqlLogAdapter.NewZapAdapter(logger)

	db = sqldblogger.OpenDriver(
		dsn,
		db.Driver(),
		loggerAdapter,
		sqldblogger.WithSQLQueryAsMessage(true),
		sqldblogger.WithPreparerLevel(sqldblogger.LevelDebug),
	)

	//client, err := sqlx.Open("mysql", "root:1234@tcp(localhost:3306)/banking")
	//if err != nil {
	//	panic(err)
	//}

	client := sqlx.NewDb(db, "mysql")

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
