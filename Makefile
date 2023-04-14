mkfile_dir:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

build:
	go build cli/iris/main.go

upgrade:
	curl https://app.vikin.gr/api/openapi.json > $(mkfile_dir)/scripts/openapi.json 
	python $(mkfile_dir)/scripts/subset_maker.py $(mkfile_dir)/scripts/openapi.json > $(mkfile_dir)/scripts/openapi-subset.json
	oapi-codegen --package=openapi -generate=types,client -o $(mkfile_dir)/pkg/openapi/openapi.gen.go $(mkfile_dir)/scripts/openapi-subset.json
	rm $(mkfile_dir)/scripts/openapi.json
	rm $(mkfile_dir)/scripts/openapi-subset.json

release:
	git tag $(RELEASE)
	git push --tags