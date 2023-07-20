SHELL := /bin/bash

wire:
	cd internal/di && wire

run:
	go run main.go

build:
	go build main.go
