package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/xuri/excelize/v2"
	"log"
	"net/http"
	"strconv"
	"web16/model"
)

var rd *render.Render = render.New()

type Success struct {
	Success bool `json:"success"`
}

type AppHandler struct {
	http.Handler
	db model.DBHandler
}

func (a *AppHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func (a *AppHandler) getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	list := a.db.GetTodos()
	rd.JSON(w, http.StatusOK, list)
}

func (a *AppHandler) addTodoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	todo := a.db.AddTodo(name)
	rd.JSON(w, http.StatusCreated, todo)
}

func (a *AppHandler) removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	ok := a.db.RemoveTodo(id)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}

}

func (a *AppHandler) completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	requestTodo := new(model.Todo)
	json.NewDecoder(r.Body).Decode(requestTodo)
	id := requestTodo.ID
	completed := requestTodo.Completed

	a.db.CompleteTodo(id, completed)

}

func (a *AppHandler) Close() {
	a.db.Close()
}

func MakeHandler(filepath string) *AppHandler {
	r := mux.NewRouter()
	a := &AppHandler{
		Handler: r,
		db:      model.NewDBHandler(filepath),
	}

	r.HandleFunc("/", a.indexHandler)
	r.HandleFunc("/todos", a.getTodoListHandler).Methods("GET")
	r.HandleFunc("/todos", a.addTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", a.removeTodoHandler).Methods("DELETE")
	r.HandleFunc("/todos", a.completeTodoHandler).Methods("PUT")
	r.HandleFunc("/excel", downloadExcelHandler)

	return a
}


func downloadExcelHandler(w http.ResponseWriter, r *http.Request) {
	file := excelize.NewFile()

	sheet := file.NewSheet("Sheet1")
	sumSheet := file.NewSheet("sum-sheet")

	file.SetActiveSheet(sheet)
	file.SetActiveSheet(sumSheet)

	file.SetCellValue("Sheet1", "A1", "VALUE")

	file.SetCellValue("sum-sheet", "A1", "SUM")
	file.SetCellValue("sum-sheet", "A2", 100)
	file.SetCellValue("sum-sheet", "B2", 250)
	file.SetCellFormula("sum-sheet", "B1", "SUM('sum-sheet'!A2,'sum-sheet'!B2)")

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment;filename=userInputData.xlsx")
	w.Header().Set("File-Name", "userInputData.xlsx")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	err := file.Write(w)
	if err != nil {
		fmt.Println("에러 발생")
		log.Fatalln("에러발생")
	}
}
