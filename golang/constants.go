package golang

var makefileText = `.DEFAULT_GOAL := test

VERSION := $(shell git describe --always --long)

help:          ## Show available options with this Makefile
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY : test
test:          ## Run all the tests
test:
	chmod +x ./test.sh && ./test.sh

clean:         ## Clean the application
clean:
	@go clean -i ./...

release:       ## Create a release build
release:	clean
	@GOOS=linux go build -i -v -ldflags="-s -w -X main.version=${VERSION}" github.com/%s/%s


ensure:        ## Run dep ensure.
ensure:
	@ dep ensure`

var readmeText = `%s:
	---

	This project can be used to ...

	Install:
	---
	Clone the project and run ...

	Test:
	---
	To run the tests: ...


	Usage:
	---

	....

	Example:
	---

	...
	`

var mainText = `package main

import(
	"fmt"
)
var version = "undefined"

func main(){
	fmt.Println("Current version is: %s", version)
}

`
