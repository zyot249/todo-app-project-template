package main

import (
	todo "todo-app/internal/todo/ports"
	"todo-app/pkg/core/http/server"
)

func main() {
	server.Start(todo.NewTodoServer())
}
