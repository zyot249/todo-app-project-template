package db

import (
	"context"
	"database/sql"
	"todo-app/internal/domain/todo"
)

type TodoRepository struct {
	db      *sql.DB
	queries *Queries
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{db: db, queries: New(db)}
}

func (r TodoRepository) Create(todo todo.Todo) (int32, error) {
	t, err := r.queries.CreateTodo(context.Background(), CreateTodoParams{
		Title:       todo.Title,
		Description: todo.Description,
		CreatedAt:   todo.CreatedAt,
	})
	if err != nil {
		return -1, err
	}
	return t.ID, nil
}

func (r TodoRepository) All() ([]todo.Todo, error) {
	todos, err := r.queries.ListTodos(context.Background())
	if err != nil {
		return nil, err
	}

	todoList := make([]todo.Todo, len(todos))
	for i, t := range todos {
		todoList[i] = todo.Todo{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			CreatedAt:   t.CreatedAt,
			Finished:    t.Finished.Bool,
		}
	}
	return todoList, nil
}

func (r TodoRepository) Get(id int32) (todo.Todo, error) {
	t, err := r.queries.GetTodo(context.Background(), id)
	if err != nil {
		return todo.Todo{}, err
	}
	return todo.Todo{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		CreatedAt:   t.CreatedAt,
		Finished:    t.Finished.Bool,
	}, nil
}

func (r TodoRepository) Update(todo todo.Todo) error {
	panic("not implemented") // TODO: Implement
}

func (r TodoRepository) Delete(id int32) error {
	err := r.queries.DeleteTodo(context.Background(), id)
	if err != nil {
		return err
	}
	return nil
}
