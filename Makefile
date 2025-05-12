include .env
export

SQLC = go run -modfile=./tools/go.mod github.com/sqlc-dev/sqlc/cmd/sqlc

.PHONY: openapi_http
openapi_http:
	@./scripts/openapi-http.sh todo internal/todo/ports todo
	@./scripts/openapi-http.sh auth internal/auth/ports auth

sqlc:
	$(SQLC) generate

run-todo:
	go run cmd/todo/main.go

run-auth:
	go run cmd/auth/main.go

run-swaggerui:
	go run cmd/swaggerui/main.go

test:
	go test -v ./...
