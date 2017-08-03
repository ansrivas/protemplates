.DEFAULT_GOAL := test

VERSION := $(shell git describe --always --long)

help:          ## Show available options with this Makefile
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY : test
test:          ## Run all the tests
test:
	./test.sh

clean:         ## Clean the application
clean:
	@go clean -i ./...
	@rm -rf ./protemplates

release:       ## Create a release build
release:	clean
	@GOOS=linux go build -o prebuilt/protemplates -i -v -ldflags="-s -w -X main.version=${VERSION}" github.com/ansrivas/protemplates


ensure:        ## Run dep ensure.
ensure:
	@ dep ensure
