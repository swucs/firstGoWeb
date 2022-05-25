package authorization

import (
	"errors"
	"fmt"
	"github.com/casbin/casbin"
	"log"
	"net/http"
)

// Authorizer is a middleware for authorization
func Authorizer(e *casbin.Enforcer, role string, uid string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if role == "" {
				role = "anonymous"
			}

			// casbin enforce
			res, err := e.EnforceSafe(role, r.URL.Path, r.Method)
			if err != nil {
				writeError(http.StatusInternalServerError, "ERROR", w, err)
				return
			}
			if res {
				next.ServeHTTP(w, r)
			} else {
				fmt.Printf("권한없음[role : %s, uid : %s]: %s\n", role, uid, r.URL.Path)
				writeError(http.StatusForbidden, "FORBIDDEN", w, errors.New("unauthorized"))
				return
			}
		}

		return http.HandlerFunc(fn)
	}
}

func writeError(status int, message string, w http.ResponseWriter, err error) {
	log.Print("ERROR: ", err.Error())
	w.WriteHeader(status)
	w.Write([]byte(message))
}
