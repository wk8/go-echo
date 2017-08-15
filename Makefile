SHELL := /bin/bash

.PHONY: run
run:
	go run echo.go

.PHONY: build
build:
	go build echo.go
