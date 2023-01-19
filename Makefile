run: ## Run the application
		GO111MODULE=on go run -mod=vendor ./cmd/main.go

test: ## Run unit tests
	go test -mod=vendor `go list ./... | grep -v 'docs'` -race