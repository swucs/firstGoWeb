package main

import (
	"casbin-test/authorization"
	"casbin-test/model"
	"github.com/casbin/casbin"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	// setup casbin auth rules
	authEnforcer, err := casbin.NewEnforcerSafe("./auth_model.conf", "./policy.csv")
	if err != nil {
		log.Fatal(err)
	}

	// setup routes
	router := mux.NewRouter()
	//mux := http.NewServeMux()

	router.HandleFunc("/", homeHandler).Methods(http.MethodGet)
	router.HandleFunc("/login", loginHandler).Methods(http.MethodGet)
	router.HandleFunc("/member/current", currentMemberHandler).Methods(http.MethodGet)
	router.HandleFunc("/member/role", memberRoleHandler).Methods(http.MethodGet)
	router.HandleFunc("/admin/stuff", adminHandler).Methods(http.MethodGet)

	users := createUsers()
	log.Print("Server started on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", authorization.Authorizer(authEnforcer, users, "admin", "swucs")(router)))
}

func createUsers() model.Users {
	users := model.Users{}
	users = append(users, model.User{ID: 1, Name: "Admin", Role: "admin"})
	users = append(users, model.User{ID: 2, Name: "Sabine", Role: "member"})
	users = append(users, model.User{ID: 3, Name: "Sepp", Role: "member"})
	return users
}

func currentMemberHandler(w http.ResponseWriter, r *http.Request) {
	writeSuccess("currentMemberHandler", w)
}

func memberRoleHandler(w http.ResponseWriter, r *http.Request) {
	writeSuccess("memberRoleHandler", w)
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	writeSuccess("adminHandler", w)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	writeSuccess("loginHandler : "+name, w)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	writeSuccess("homeHandler", w)
}

func writeError(status int, message string, w http.ResponseWriter, err error) {
	log.Print("ERROR: ", err.Error())
	w.WriteHeader(status)
	w.Write([]byte(message))
}

func writeSuccess(message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}
