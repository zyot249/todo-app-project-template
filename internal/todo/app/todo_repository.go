package app

import "todo-app/internal/todo/domain"

type TodoRepository interface {
	Create(todo domain.Todo) (int32, error)
	All() ([]domain.Todo, error)
	Get(id int32) (domain.Todo, error)
	Update(todo domain.Todo) error
	Delete(id int32) error
}
