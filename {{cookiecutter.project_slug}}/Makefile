fmt: ## Format code
	gofumpt -l -w $$(go list -f {{.Dir}} ./... | grep -v /vendor/)

lint: ## Run lint go code
	golangci-lint run