package todo

import (
	"net/http"
	"todo-app/internal/common/core/logs"
	"todo-app/internal/common/core/sql"
	db "todo-app/internal/todo/adapters"
	"todo-app/internal/todo/app"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type TodoHTTPServer struct {
	app app.Application
}

func NewTodoHTTPServer() TodoHTTPServer {
	logs.Init()

	logger := logrus.NewEntry(logrus.StandardLogger())

	sqlDbConfig := sql.NewConfig()
	sqlDb, err := sql.NewDBConnection(sqlDbConfig)
	if err != nil {
		logger.Fatalf("failed to initialize database: %v", err)
	}

	app := app.Application{
		Logger:         logger,
		TodoRepository: db.NewPostgresTodoRepository(sqlDb),
	}
	return TodoHTTPServer{app: app}
}

func (s TodoHTTPServer) GetPort() string {
	return "8080"
}

func (s TodoHTTPServer) GetServiceName() string {
	return "todo"
}

func (s TodoHTTPServer) RegisterMiddlewares(router *chi.Mux) {
}

func (s TodoHTTPServer) CreateApiHandler(router *chi.Mux) http.Handler {
	return HandlerFromMux(NewTodoHandler(s.app), router)
}

func (s TodoHTTPServer) CreateRootHandler(router *chi.Mux) http.Handler {
	return router
}

func (s TodoHTTPServer) HaveApiSwaggerDoc() bool {
	return true
}
