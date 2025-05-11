package todo

import "time"

type Todo struct {
	ID          int32
	Title       string
	Description string
	CreatedAt   time.Time
	Finished    bool
}

func NewTodo(title, description string) *Todo {
	return &Todo{
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
		Finished:    false,
	}
}

func (t Todo) IsFinished() bool {
	return t.Finished
}

func (t *Todo) Finish() {
	t.Finished = true
}
