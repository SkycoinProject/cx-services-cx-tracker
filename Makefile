lint: ## Run linters.
		GO111MODULE=off vendorcheck ./...
		golangci-lint run -c .golangci.yml ./...

install-linters: ## Install linters
		go get -u github.com/FiloSottile/vendorcheck
		go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

format: ## Formats the code. Must have goimports installed
		goimports -w -local github.com/watercompany/cx-tracker ./cmd
		goimports -w -local github.com/watercompany/cx-tracker ./src

test:
		@mkdir -p coverage/
		go test -race -tags no_ci -cover -timeout=5m ./...