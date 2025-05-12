package main

import (
	"todo-app/internal/common/core/http/server"
	"todo-app/internal/common/core/logs"
	"todo-app/internal/swaggerui"
)

func main() {
	logs.Init()

	server.Start(swaggerui.NewSwaggerUIServer())
}
