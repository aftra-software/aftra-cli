#!/usr/bin/env bash
set -euxo pipefail
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
MODULE_OUT="$SCRIPT_DIR"

openapi_file=$(realpath "$SCRIPT_DIR/openapi.json")


docker run --rm \
-v "$SCRIPT_DIR:$SCRIPT_DIR" \
-v "$MODULE_OUT:/out" \
--user "$(id -u):$(id -g)" \
openapitools/openapi-generator-cli:v6.2.0 generate \
-i "$openapi_file" \
-g go \
-o /out/ \
--skip-validate-spec
