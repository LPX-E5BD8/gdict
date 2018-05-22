# This how we want to name the binary output
BINARY=gdict
M = $(shell printf "\033[34;1m▶\033[0m")

all: fmt build

# Code format
.PHONY: fmt
fmt: ; $(info $(M) running gofmt…) @ ## Run gofmt on all source files
	@ret=0 && for d in $$(go list -f '{{.Dir}}' ./... | grep -v /vendor/); do \
		gofmt -l -w $$d/*.go || ret=$$? ; \
	done ; exit $$ret

# Builds the project
build:
	go build

# Installs our project: copies binaries
install:
	go install