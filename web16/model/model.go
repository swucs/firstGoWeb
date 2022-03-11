package model

import "time"

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

var todoMap map[int]*Todo

func GetTodos() []*Todo {
	return nil
}

func AddTodo(name string) *Todo {
	return nil
}

func RemoveTodo(id int) bool {
	return true
}
