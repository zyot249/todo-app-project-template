package main

import (
	"todo-app/internal/common/core/http/server"
	todo "todo-app/internal/todo/ports"
)

func main() {
	server.Start(todo.NewTodoServer())
}
