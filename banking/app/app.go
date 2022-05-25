package app

import (
	"banking/authorization"
	"banking/config"
	module "banking/wire"

	"fmt"
	"github.com/casbin/casbin"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
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

	//Database 커넥션 초기화
	config.InitDatabaseConnection()
	client := config.GetDbClient()
	if client == nil {
		panic("database is not connected.")
	}

	//wire 로 DI 주입
	customerHandler := module.InitializeCustomer()
	accountHandler := module.InitializeAccount()

	router.HandleFunc("/customers", customerHandler.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", accountHandler.NewAccount).Methods(http.MethodPost)

	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	//권한체크 설정
	enforcer := createAuthEnforcer()

	//서버시작
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), authorization.Authorizer(enforcer, "member", "test")(router)))
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")
}
