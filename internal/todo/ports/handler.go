package todo

import (
	"encoding/json"
	"net/http"
	"todo-app/internal/todo/app"
	"todo-app/internal/todo/constants"
	"todo-app/internal/todo/domain"
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
func (h TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request, id int) {
	err := h.app.TodoRepository.Delete(int32(id))
	if err != nil {
		response.RespondError(w, r, constants.TodoMessageIdConstants.DeleteTodo(), errors.FromError(err))
		return
	}
	response.RespondSuccess(w, r, constants.TodoMessageIdConstants.DeleteTodo(), nil)
}

// (GET /todo/{id})
func (h TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request, id int) {
	todo, err := h.app.TodoRepository.Get(int32(id))
	if err != nil {
		response.RespondError(w, r, constants.TodoMessageIdConstants.GetTodo(), errors.FromError(err))
		return
	}
	response.RespondSuccess(w, r, constants.TodoMessageIdConstants.GetTodo(), todo)
}

// (PUT /todo/{id})
func (h TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request, id int) {
	panic("not implemented") // TODO: Implement
}

// (GET /todos)
func (h TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.app.TodoRepository.All()
	if err != nil {
		response.RespondError(w, r, constants.TodoMessageIdConstants.GetTodo(), errors.FromError(err))
		return
	}
	response.RespondSuccess(w, r, constants.TodoMessageIdConstants.GetTodo(), todos)
}

// (POST /todos)
func (h TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var createTodoDto CreateTodoDto
	err := json.NewDecoder(r.Body).Decode(&createTodoDto)
	if err != nil {
		response.RespondError(w, r, constants.TodoMessageIdConstants.CreateTodo(), errors.FromError(err))
		return
	}

	todo, err := h.app.TodoRepository.Create(domain.Todo{
		Title:       createTodoDto.Title,
		Description: createTodoDto.Description,
	})

	if err != nil {
		response.RespondError(w, r, constants.TodoMessageIdConstants.CreateTodo(), errors.FromError(err))
		return
	}

	response.RespondSuccess(w, r, constants.TodoMessageIdConstants.CreateTodo(), todo)
}
