#!/bin/bash
set -e

readonly service="$1"
readonly output_dir="$2"
readonly package="$3"

OAPI_CODEGEN="go run -modfile=./tools/go.mod github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"

$OAPI_CODEGEN -generate types -o "$output_dir/openapi_types.gen.go" -package "$service" "api/openapi/$service.yml"
$OAPI_CODEGEN -generate chi-server -o "$output_dir/openapi_server.gen.go" -package "$service" "api/openapi/$service.yml"
# $OAPI_CODEGEN -generate types -o "internal/common/client/$service/openapi_types.gen.go" -package "$service" "api/openapi/$service.yml"
# $OAPI_CODEGEN -generate client -o "internal/common/client/$service/openapi_client.gen.go" -package "$service" "api/openapi/$service.yml"
