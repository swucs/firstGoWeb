package config

import (
	"banking/sqlLogAdapter"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var dbClient *sqlx.DB

func GetDbClient() *sqlx.DB {
	return dbClient
}

/**
Database 커넥션 초기화 (mysql, sqlx)
*/
func InitDatabaseConnection() {

	//go 실행시 환경변수 입력
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	//root:1234@tcp(localhost:3306)/banking
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	fmt.Printf("####### datasource : %s\n", dataSource)
	db, err := sql.Open("mysql", dataSource)
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
		dataSource,
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

	if client == nil {
		fmt.Println("Client is null")
	} else {
		fmt.Println("Client is connected........")
	}

	dbClient = client
}
