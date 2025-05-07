package main

import (
	"todo-app/internal/ports/swaggerui"
	"todo-app/pkg/core/http/server"
	"todo-app/pkg/core/logs"
)

func main() {
	logs.Init()

	server.Start(swaggerui.NewSwaggerUIServer())
}
