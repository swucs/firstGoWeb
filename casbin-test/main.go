package main

import (
	"casbin-test/authorization"
	"casbin-test/model"
	"github.com/casbin/casbin"
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
	mux := http.NewServeMux()

	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/member/current", currentMemberHandler)
	mux.HandleFunc("/member/role", memberRoleHandler)
	mux.HandleFunc("/admin/stuff", adminHandler)

	users := createUsers()
	log.Print("Server started on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", authorization.Authorizer(authEnforcer, users, "admin", "swucs")(mux)))
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
	name := r.PostFormValue("name")
	writeSuccess("loginHandler : "+name, w)
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
