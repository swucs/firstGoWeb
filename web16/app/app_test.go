package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"testing"
	"web16/model"
)

func TestTodos(t *testing.T) {
	filepath := "./test.db"
	os.Remove(filepath)

	assert := assert.New(t)

	ah := MakeHandler(filepath)
	defer ah.Close()

	ts := httptest.NewServer(ah)
	defer ts.Close()

	fmt.Println("test")

	resp, err := http.PostForm(ts.URL+"/todos", url.Values{"name": {"test todo"}})
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	var todo model.Todo
	err = json.NewDecoder(resp.Body).Decode(&todo)
	assert.NoError(err)
	assert.Equal("test todo", todo.Name)
	id1 := todo.ID

	resp, err = http.PostForm(ts.URL+"/todos", url.Values{"name": {"test todo2"}})
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	err = json.NewDecoder(resp.Body).Decode(&todo)
	assert.NoError(err)
	assert.Equal(todo.Name, "test todo2")
	id2 := todo.ID

	resp, err = http.Get(ts.URL + "/todos")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	todos := []*model.Todo{}
	err = json.NewDecoder(resp.Body).Decode(&todos)
	assert.NoError(err)
	assert.Equal(2, len(todos))

	for _, t := range todos {
		if t.ID == id1 {
			assert.Equal("test todo", t.Name)
		} else if t.ID == id2 {
			assert.Equal("test todo2", t.Name)
		} else {
			assert.Fail("fail")
		}
	}

	reqTodo := &model.Todo{ID: id2, Completed: true}
	data, _ := json.Marshal(reqTodo)

	req, err := http.NewRequest("PUT", ts.URL+"/todos", bytes.NewBuffer(data))
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	req, _ = http.NewRequest("DELETE", ts.URL+"/todos/"+strconv.Itoa(id1), nil)
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	resp, err = http.Get(ts.URL + "/todos")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	todos = []*model.Todo{}
	err = json.NewDecoder(resp.Body).Decode(&todos)
	assert.NoError(err)
	assert.Equal(1, len(todos))

	for _, t := range todos {
		assert.Equal(t.ID, id2)
	}
}
