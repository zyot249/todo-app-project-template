include .env
export

OAPI_CODEGEN = go run -modfile=./tools/go.mod github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen
OAPI_DIRS := $(wildcard internal/ports/http/*)

openapi:
	@for dir in $(OAPI_DIRS); do \
		if [ -d $$dir ]; then \
			echo "Processing $$dir"; \
			cd $$dir && \
			go run -modfile=../../../../tools/go.mod github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen \
			-generate chi-server \
			-config ./openapi.cfg.yml \
			-o ./openapi_server.gen.go \
			./openapi.yml; \
			go run -modfile=../../../../tools/go.mod github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen \
			-generate types \
			-config ./openapi.cfg.yml \
			-o ./openapi_types.gen.go \
			./openapi.yml; \
			cd - > /dev/null; \
		fi \
	done

run:
	go run cmd/todo/main.go

test:
	go test -v ./...
