build:
	go build cli/iris/main.go
	
upgrade:
	python scripts/subset_maker.py scripts/openapi.json > scripts/openapi-subset.json
	oapi-codegen --package=openapi -generate=types -o ./pkg/openapi/openapi.gen.go scripts/openapi-subset.json
