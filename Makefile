run: ## Run the application
		GO111MODULE=on go run -mod=vendor ./cmd/main.go

test: ## Run unit tests
	go test -mod=vendor `go list ./... | grep -v 'docs'` -race

lint: ## Perform linting
	golangci-lint run --disable-all -E revive  --exclude-use-default=false --modules-download-mode=vendor