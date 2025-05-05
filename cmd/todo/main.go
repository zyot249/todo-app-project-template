package main

import (
	"net/http"
	"todo-app/internal/app"
	"todo-app/internal/ports/http/todo"
	"todo-app/pkg/core/logs"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	logs.Init()

	logger := logrus.NewEntry(logrus.StandardLogger())

	app := app.NewApplication(logger)

	todo.RunServer(func(router chi.Router) http.Handler {
		return todo.HandlerFromMux(todo.NewTodoHandler(app), router)
	})
}
