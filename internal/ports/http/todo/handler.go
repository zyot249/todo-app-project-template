package todo

import (
	"net/http"
	"time"
	"todo-app/internal/app"
	"todo-app/internal/ports/constants"
	"todo-app/pkg/core/errors"
	"todo-app/pkg/core/http/response"
)

type TodoHandler struct {
	app app.Application
}

func NewTodoHandler(app app.Application) TodoHandler {
	return TodoHandler{app: app}
}

// (DELETE /todo/{id})
func (h TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request, id string) {
	panic("not implemented") // TODO: Implement
}

// (GET /todo/{id})
func (h TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request, id string) {
	response.RespondSuccess(w, r, constants.TodoMessageIdConstants.GetTodo(), Todo{
		Id:            id,
		Title:         "Test",
		Completed:     false,
		CreatedAt:     time.Now(),
		LastUpdatedAt: time.Now(),
	})
}

// (PUT /todo/{id})
func (h TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request, id string) {
	panic("not implemented") // TODO: Implement
}

// (GET /todos)
func (h TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	response.RespondError(w, r, constants.TodoMessageIdConstants.GetTodo(), errors.SimpleShynobiError{
		Message: "test",
		Code:    constants.TodoErrorCodeConstants.Unimplemented(),
	})
}

// (POST /todos)
func (h TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	panic("not implemented") // TODO: Implement
}
