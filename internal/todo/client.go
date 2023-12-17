package todo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const API_HOST = "https://jsonplaceholder.typicode.com"

type TodoClient struct {
	Client *http.Client
}

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func NewTodoClient() *TodoClient {
	client := &http.Client{}
	return &TodoClient{Client: client}
}

func (c *TodoClient) FetchTodo(id int) (Todo, error) {
	url := fmt.Sprintf("%s/todos/%d", API_HOST, id)
	resp, err := c.Client.Get(url)
	if err != nil {
		return Todo{}, fmt.Errorf("error fetching TODO %d: %v", id, err)
	}
	defer resp.Body.Close()

	var todo Todo
	if err := DecodeResponse(resp.Body, &todo); err != nil {
		return Todo{}, fmt.Errorf("error decoding TODO %d response: %v", id, err)
	}

	return todo, nil
}

func DecodeResponse(body io.ReadCloser, v interface{}) error {
	defer body.Close()
	decoder := json.NewDecoder(body)
	return decoder.Decode(v)
}
