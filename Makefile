ROOT_APP_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: all
all: run

.PHONY: build
build: ## Build app
	CGO_ENABLED=0 go build -o bin/ivy cmd/ivy/main.go

.PHONY: build-race
build-race: ## Build app with race flag
	CGO_ENABLED=1 go build -race -o bin/ivyr cmd/ivy/main.go

.PHONY: run
run: build ## Build and run with default config
	bin/ivy --config=configs/ivy-default.conf.yml

.PHONY: clean
clean: ## Clean bin dir
	rm -rf bin/

.PHONY: run-race
run-race: build-race ## Build with race flag and run with default config
	bin/ivyr --config=configs/ivy-default.conf.yml

.PHONY: fmt
fmt: ## Run go fmt and goimports
	go fmt ./...
	goimports -w ./

.PHONY: lint
lint: ## Run golangci-lint
	golangci-lint run -v ./...

.PHONY: test
test: ## Run tests
	@go test ./internal...

.PHONY: cover
cover: ## Run tests with cover
	@go test -coverprofile=coverage.out ./internal...

.PHONY: coverr
coverr: cover ## Run tests with cover and open report
	@go tool cover -html=coverage.out

.PHONY: help
help: ## List of commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'