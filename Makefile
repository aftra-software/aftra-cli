mkfile_dir:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

OPENAPI_HOST ?= https://app.aftra.io

build:
	go build cli/aftra/main.go

upgrade:
	echo ${OPENAPI_HOST}
	curl ${OPENAPI_HOST}/api/openapi.json > $(mkfile_dir)/scripts/openapi.json 
	python $(mkfile_dir)/scripts/subset_maker.py $(mkfile_dir)/scripts/openapi.json > $(mkfile_dir)/scripts/openapi-subset.json
	oapi-codegen --package=openapi -generate=types,client -o $(mkfile_dir)/pkg/openapi/openapi.gen.go $(mkfile_dir)/scripts/openapi-subset.json
	rm $(mkfile_dir)/scripts/openapi.json
	rm $(mkfile_dir)/scripts/openapi-subset.json

release:
	git tag $(RELEASE)
	git push --tags