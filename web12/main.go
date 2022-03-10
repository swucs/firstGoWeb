package main

import (
	"encoding/json"
	"github.com/gorilla/pat"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"net/http"
	"time"
)

var rd *render.Render

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "tucker", Email: "tucker@naver.com", CreatedAt: time.Now()}

	rd.JSON(w, http.StatusOK, user)
	//
	//w.Header().Add("Content-type", "application/json")
	//w.WriteHeader(http.StatusOK)
	//data, _ := json.Marshal(user)
	//fmt.Fprintf(w, string(data))
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		//w.WriteHeader(http.StatusBadRequest)
		//fmt.Fprint(w, err)
		rd.Text(w, http.StatusBadRequest, err.Error())
		return
	}

	user.CreatedAt = time.Now()
	rd.JSON(w, http.StatusOK, user)

	//w.Header().Add("Content-type", "application/json")
	//w.WriteHeader(http.StatusOK)
	//data, _ := json.Marshal(user)
	//fmt.Fprintf(w, string(data))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "tucker", Email: "tucker@naver.com", CreatedAt: time.Now()}
	rd.HTML(w, http.StatusOK, "body", user) //"body"는 tmpl확장자명은 뺀다.
}

func main() {
	rd = render.New(render.Options{
		Directory:  "view",                     //directory 지정(원래 /templates 임)
		Extensions: []string{".html", ".tmpl"}, //template 확장자명 지정 (원래 확장자명은 "tmpl"으로 고정되어 있음)
		Layout:     "hello",
	})
	mux := pat.New()

	mux.Get("/users", getUserInfoHandler)
	mux.Post("/users", addUserHandler)
	mux.Get("/hello", helloHandler)

	//negroni (public 디렉토리 아래를 실행시킴)
	n := negroni.Classic()
	n.UseHandler(mux)
	http.ListenAndServe(":3000", n)
}
