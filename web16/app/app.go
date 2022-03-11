package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
	"strconv"
	"web16/model"
)

var rd *render.Render

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	//list := []*model.Todo{}
	//for _, v := range todoMap {
	//	list = append(list, v)
	//}
	list := model.GetTodos()
	rd.JSON(w, http.StatusOK, list)
}

func addTodoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	todo := model.AddTodo(name)
	//id := len(todoMap) + 1
	//todo := &Todo{id, name, false, time.Now()}
	//todoMap[id] = todo
	rd.JSON(w, http.StatusCreated, todo)
}

//func addTestTodos() {
//	todoMap[1] = &Todo{1, "Buy a milk", false, time.Now()}
//	todoMap[2] = &Todo{2, "Exercise", true, time.Now()}
//	todoMap[3] = &Todo{3, "Home work", false, time.Now()}
//}

type Success struct {
	Success bool `json:"success"`
}

func removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	ok := model.RemoveTodo(id)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
	//if _, ok := todoMap[id]; ok {
	//	delete(todoMap, id)
	//	rd.JSON(w, http.StatusOK, Success{true})
	//} else {
	//	rd.JSON(w, http.StatusOK, Success{false})
	//}
}

func completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	requestTodo := new(model.Todo)
	json.NewDecoder(r.Body).Decode(requestTodo)
	id := requestTodo.ID
	completed := requestTodo.Completed

	model.CompleteTodo(id, )

	//if todo, ok := todoMap[id]; ok {
	//	todo.Completed = completed
	//	rd.JSON(w, http.StatusOK, Success{true})
	//} else {
	//	rd.JSON(w, http.StatusOK, Success{false})
	//}
}

func MakeHandler() http.Handler {
	todoMap = make(map[int]*Todo)
	//addTestTodos()
	rd = render.New()
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/todos", getTodoListHandler).Methods("GET")
	r.HandleFunc("/todos", addTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", removeTodoHandler).Methods("DELETE")
	r.HandleFunc("/todos", completeTodoHandler).Methods("PUT")

	return r
}
