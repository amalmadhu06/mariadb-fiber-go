
wire: ## Generate wire_gen.go
	cd pkg/di && wire

run:
	cd cmd/api && go run main.go