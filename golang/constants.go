package golang

var makefileText = `.DEFAULT_GOAL := help

DEP= $(shell command -v dep 2>/dev/null)
VERSION=$(shell git describe --always --long)
PROJECT_NAME := %s
CLONE_URL:=github.com/%s/%s
IDENTIFIER= $(VERSION)-$(GOOS)-$(GOARCH)
BUILD_TIME=$(shell date -u +%FT%T%z)
LDFLAGS="-s -w -X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)"

help:          ## Show available options with this Makefile
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY : test crossbuild release
test:          ## Run all the tests
test:
	chmod +x ./test.sh && ./test.sh

clean:         ## Clean the application
clean:
	@go clean -i ./...
	@rm -rf ./{PROJECT_NAME}

build: vendor	clean
	go build -i -v -ldflags $(LDFLAGS) $(FLAGS) $(CLONE_URL)

dep:           ## Go get dep
dep:
	go get -u github.com/golang/dep/cmd/dep

ensure:        ## Run dep ensure.
ensure:
ifndef DEP
	make dep
endif
	dep ensure
	touch vendor

crossbuild: ensure
	mkdir -p build/${PROJECT_NAME}-$(IDENTIFIER)
	make build FLAGS="-o build/${PROJECT_NAME}-$(IDENTIFIER)/${PROJECT_NAME}"

release:       ## Create a release build.
release:
	make crossbuild GOOS=linux GOARCH=amd64
	make crossbuild GOOS=linux GOARCH=386
	make crossbuild GOOS=darwin GOARCH=amd64


bench:	       ## Benchmark the code.
bench:
	@go test -o bench.test -cpuprofile cpu.prof -memprofile mem.prof -bench .

prof:          ## Run the profiler.
prof:	bench
	@go tool pprof cpu.prof

prof_svg:      ## Run the profiler and generate image.
prof_svg:	clean	bench
	@echo "Do you have graphviz installed? sudo apt-get install graphviz."
	@go tool pprof -svg bench.test cpu.prof > cpu.svg`

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

import (
	"fmt"
)

// Version is used to set the current version of this build.
var Version = "undefined"

func main() {
	fmt.Printf("Current version is: %s", Version)
}
`
