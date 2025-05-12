package app

import (
	"github.com/sirupsen/logrus"
)

type Application struct {
	Logger         *logrus.Entry
	TodoRepository TodoRepository
}
