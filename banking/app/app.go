package app

import (
	"banking/authorization"
	"banking/domain"
	"banking/service"
	"banking/sqlLogAdapter"
	"database/sql"
	"fmt"
	"github.com/casbin/casbin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"time"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variable not defined...")
	}
}

/**

 */
func createAuthEnforcer() *casbin.Enforcer {

	//if os.Getenv("AUTH_MODEL_FILEPATH") == "" ||
	//	os.Getenv("AUTH_POLICY_FILEPATH") == "" {
	//	log.Fatal("The Casbin environment variable not defined...")
	//}

	//casbin 권한정책 세팅
	authEnforcer, err := casbin.NewEnforcerSafe("./auth_model.conf", "./policy.csv")
	if err != nil {
		log.Fatal(err)
	}

	return authEnforcer
}

func Start() {

	sanityCheck()

	//mux := http.NewServeMux()
	router := mux.NewRouter()

	dbClient := getDbClient()

	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.newAccount).Methods(http.MethodPost)

	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	//권한체크 설정
	enforcer := createAuthEnforcer()

	//서버시작
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), authorization.Authorizer(enforcer, "admin", "test")(router)))
}

func getDbClient() *sqlx.DB {

	//go 실행시 환경변수 입력
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	//root:1234@tcp(localhost:3306)/banking
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
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

	return client
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")
}
