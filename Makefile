openapi-generator-go generate -i api/openapi.json -g go -o /tmp/test/

openapi-generator-go generate models --spec ./example/api.yaml --output ./example/models --package-name models
