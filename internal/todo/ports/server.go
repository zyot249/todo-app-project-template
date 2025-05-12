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

type TodoServer struct {
	app app.Application
}

func NewTodoServer() TodoServer {
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
