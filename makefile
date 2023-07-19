
wire: ## Generate wire_gen.go
	cd internal/di && wire

run:
	cd cmd/api && go run main.go