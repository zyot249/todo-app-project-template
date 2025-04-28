OAPI_CODEGEN = go run -modfile=./tools/go.mod github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen
OAPI_DIRS := $(wildcard internal/ports/http/*)

.PHONY: gen_openapi

gen_openapi:
	@for dir in $(OAPI_DIRS); do \
		if [ -d $$dir ]; then \
			echo "Processing $$dir"; \
			cd $$dir && \
			go run -modfile=../../../../tools/go.mod github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen \
			-config ./openapi.cfg.yml \
			-o ./openapi.gen.go \
			./openapi.yml; \
			cd - > /dev/null; \
		fi \
	done

.PHONY: run
run:
	go run cmd/todoapp/main.go

.PHONY: test
test:
	go test -v ./...
