package main

import (
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"web16/app"
)

func main() {

	m := app.MakeHandler("./test.db")
	defer m.Close()

	n := negroni.Classic()
	n.UseHandler(m)

	log.Println("Started App")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}
