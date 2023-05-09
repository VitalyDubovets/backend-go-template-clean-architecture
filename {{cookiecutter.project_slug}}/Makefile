# ==============================================================================
# Running linters
fmt: ## Format code
	gofumpt -l -w $$(go list -f {{.Dir}} ./... | grep -v /vendor/)

lint: ## Run lint go code
	golangci-lint run
# ==============================================================================
# Modules support
tidy:
	go mod tidy
	go mod vendor
# ==============================================================================
# Migrations
create-migration:
	migrate create -ext sql -dir internal/infra/db/migrations -seq $(MESSAGE)

up-migration:
	migrate -path database/migration/ -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up

down-migration:
	migrate -path database/migration/ -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose down 1

fix-migration:
	migrate -path database/migration/ -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" force $(VERSION)
# ==============================================================================
# Dependencies
install-deps:
	go install mvdan.cc/gofumpt
	go install github.com/swaggo/swag
# ==============================================================================
