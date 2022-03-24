package main

import (
	"context"
	"fmt"
	"github.com/go-session/redis/v3"
	"github.com/go-session/session/v3"
	"net/http"
)

func main() {
	session.InitManager(
		session.SetStore(redis.NewRedisStore(&redis.Options{
			Addr: "127.0.0.1:6379",
			DB:   15,
		})),
	)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		store, err := session.Start(context.Background(), w, r)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}

		store.Set("foo", "bar")
		store.Set("test", "테스트")
		err = store.Save()
		if err != nil {
			fmt.Fprint(w, err)
			return
		}

		//http.Redirect(w, r, "/foo", 302)
	})

	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		store, err := session.Start(context.Background(), w, r)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}

		foo, ok := store.Get("foo")
		if ok {
			fmt.Fprintf(w, "foo:%s\n", foo)
			//return
		}

		test, ok := store.Get("test")
		if ok {
			fmt.Fprintf(w, "test:%s\n", test)
			//return
		}

		//fmt.Fprint(w, "does not exist")
	})

	http.ListenAndServe(":8080", nil)
}
