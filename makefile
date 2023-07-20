SHELL := /bin/bash

wire:
	cd internal/di && wire

run:
	go run main.go

build:
	go build main.go

mock:
	mockgen -source=internal/services/interfaces/offer.go -destination=internal/services/mockServices/mock.go -package mockServices
	mockgen -source=internal/memory/interfaces/offer.go -destination=internal/memory/mockMemory/mock.go -package mockMemory
	mockgen -source=internal/repository/interfaces/offer.go -destination=internal/repository/mockRepo/mock.go -package mockRepo
