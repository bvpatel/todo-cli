package main

import (
	"todo-cli/cmd/todo"
)

func main() {
	app := todo.NewTodoApp()
	app.Execute()
}
