package main

import (
	"todo-app/internal/swaggerui"
	"todo-app/pkg/core/http/server"
	"todo-app/pkg/core/logs"
)

func main() {
	logs.Init()

	server.Start(swaggerui.NewSwaggerUIServer())
}
