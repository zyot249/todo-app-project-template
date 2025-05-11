package repositories

import (
	"todo-app/internal/domain/todo"
)

type ITodoRepository interface {
	Create(todo todo.Todo) (int32, error)
	All() ([]todo.Todo, error)
	Get(id int32) (todo.Todo, error)
	Update(todo todo.Todo) error
	Delete(id int32) error
}
