package app

import (
	"todo-app/internal/adapters/repositories"

	"github.com/sirupsen/logrus"
)

type Application struct {
	Logger         *logrus.Entry
	TodoRepository repositories.ITodoRepository
}
