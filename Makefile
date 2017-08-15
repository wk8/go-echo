SHELL := /bin/bash

# TODO wkpo echo?
.PHONY: run
run:
	go run main.go

.PHONY: build
build:
	go build main.go
