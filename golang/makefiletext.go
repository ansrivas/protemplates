package golang

var makefileText = `.DEFAULT_GOAL := help

DEP= $(shell command -v dep 2>/dev/null)
VERSION=$(shell git describe --always --long)
PROJECT_NAME := {{.appname}}
CLONE_URL:={{.scm}}/{{.scmusername}}/{{.appname}}
IDENTIFIER= $(VERSION)-$(GOOS)-$(GOARCH)
BUILD_TIME=$(shell date -u +%FT%T%z)
LDFLAGS='-extldflags "-static" -s -w -X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)'

help:          ## Show available options with this Makefile
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY : test crossbuild release build clean
test:          ## Run all the tests
test:
	chmod +x ./test.sh && ./test.sh

clean:         ## Clean the application
clean:
	@go clean -i ./...
	@rm -rf ./$(PROJECT_NAME)
	@rm -rf build/*

# -v so warnings from the linker aren't suppressed.
# -a so dependencies are rebuilt (they may have been dynamically
# linked).
build: vendor
	xgo -go 1.9.2 -out=$(FLAGS) -ldflags=$(LDFLAGS) -targets='${GOOS}/${GOARCH}' .

dep:           ## Go get dep
dep:
        go get -u github.com/golang/dep/cmd/dep

xgo:           ## Go get XGO
xgo:
        go get -u github.com/karalabe/xgo

ensure:        ## Run dep ensure.
ensure:
ifndef DEP
        make dep
endif
        dep ensure
        touch vendor

ifndef XGO
       make xgo
endif

crossbuild:
	mkdir -p build/${PROJECT_NAME}-$(IDENTIFIER)
	make build FLAGS="build/${PROJECT_NAME}-$(IDENTIFIER)/${PROJECT_NAME}"
	cd build \
	&& tar cvzf "${PROJECT_NAME}-$(IDENTIFIER).tgz" "${PROJECT_NAME}-$(IDENTIFIER)" \
	&& rm -rf "${PROJECT_NAME}-$(IDENTIFIER)"

release:       ## Create a release build.
release:	ensure	clean
	make crossbuild GOOS=linux GOARCH=amd64
	make crossbuild GOOS=linux GOARCH=386
	make crossbuild GOOS=darwin GOARCH=amd64
	make crossbuild GOOS=windows GOARCH=amd64

bench:	       ## Benchmark the code.
bench:
	@go test -o bench.test -cpuprofile cpu.prof -memprofile mem.prof -bench .

prof:          ## Run the profiler.
prof:	bench
	@go tool pprof cpu.prof

prof_svg:      ## Run the profiler and generate image.
prof_svg:	clean	bench
	@echo "Do you have graphviz installed? sudo apt-get install graphviz."
	@go tool pprof -svg bench.test cpu.prof > cpu.svg
`
