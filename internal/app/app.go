package app

import "github.com/sirupsen/logrus"

type Application struct {
	logger *logrus.Entry
}

func NewApplication(logger *logrus.Entry) Application {
	return Application{
		logger: logger,
	}
}
