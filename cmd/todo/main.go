package main

import (
	"todo-app/internal/ports/http/todo"
	"todo-app/pkg/core/http/server"
)

func main() {
	server.Start(todo.NewTodoServer())
}
