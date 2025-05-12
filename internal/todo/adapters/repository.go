package db

import (
	"context"
	"database/sql"
	"todo-app/internal/todo/domain"
)

type PostgresTodoRepository struct {
	db      *sql.DB
	queries *Queries
}

func NewPostgresTodoRepository(db *sql.DB) *PostgresTodoRepository {
	return &PostgresTodoRepository{db: db, queries: New(db)}
}

func (r PostgresTodoRepository) Create(todo domain.Todo) (int32, error) {
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

func (r PostgresTodoRepository) All() ([]domain.Todo, error) {
	todos, err := r.queries.ListTodos(context.Background())
	if err != nil {
		return nil, err
	}

	todoList := make([]domain.Todo, len(todos))
	for i, t := range todos {
		todoList[i] = domain.Todo{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			CreatedAt:   t.CreatedAt,
			Finished:    t.Finished.Bool,
		}
	}
	return todoList, nil
}

func (r PostgresTodoRepository) Get(id int32) (domain.Todo, error) {
	t, err := r.queries.GetTodo(context.Background(), id)
	if err != nil {
		return domain.Todo{}, err
	}
	return domain.Todo{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		CreatedAt:   t.CreatedAt,
		Finished:    t.Finished.Bool,
	}, nil
}

func (r PostgresTodoRepository) Update(todo domain.Todo) error {
	panic("not implemented") // TODO: Implement
}

func (r PostgresTodoRepository) Delete(id int32) error {
	err := r.queries.DeleteTodo(context.Background(), id)
	if err != nil {
		return err
	}
	return nil
}
