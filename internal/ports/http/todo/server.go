package todo

import (
	"net/http"
	"todo-app/internal/app"
	"todo-app/pkg/core/logs"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type TodoServer struct {
	app app.Application
}

func NewTodoServer() TodoServer {
	logs.Init()

	logger := logrus.NewEntry(logrus.StandardLogger())

	app := app.NewApplication(logger)
	return TodoServer{app: app}
}

func (s TodoServer) GetHost() string {
	return "localhost"
}

func (s TodoServer) GetPort() string {
	return "8080"
}

func (s TodoServer) GetServiceName() string {
	return "todo"
}

func (s TodoServer) RegisterMiddlewares(router *chi.Mux) {
}

func (s TodoServer) CreateApiHandler(router *chi.Mux) http.Handler {
	return HandlerFromMux(NewTodoHandler(s.app), router)
}

func (s TodoServer) CreateRootHandler(router *chi.Mux) http.Handler {
	return router
}

func (s TodoServer) HaveApiSwaggerDoc() bool {
	return true
}
