.PHONY: all build test clean docker-build docker-push

# Variables
DOCKER_REGISTRY ?= your-registry
VERSION ?= $(shell git describe --tags --always --dirty)
PLATFORMS ?= linux/arm64,linux/arm/v7,linux/amd64

# Go related variables
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin

# Build all services
all: clean build test

# Build binaries
build:
	@echo "Building services..."
	@for service in auth gpio metrics; do \
		echo "Building $$service..."; \
		cd services/$$service && go build -v -o ../../bin/$$service; \
		cd $(GOBASE); \
	done

# Run tests with coverage
test:
	@echo "Running tests..."
	@for service in auth gpio metrics; do \
		echo "Testing $$service..."; \
		cd services/$$service && go test -v -race -cover ./...; \
		cd $(GOBASE); \
	done

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf bin/
	@go clean -testcache

# Build Docker images
docker-build:
	@echo "Building Docker images..."
	docker buildx create --use --name=crossplatform --node=crossplatform
	docker buildx build --platform=$(PLATFORMS) -t $(DOCKER_REGISTRY)/auth:$(VERSION) ./services/auth --push
	docker buildx build --platform=$(PLATFORMS) -t $(DOCKER_REGISTRY)/gpio:$(VERSION) ./services/gpio --push
	docker buildx build --platform=$(PLATFORMS) -t $(DOCKER_REGISTRY)/metrics:$(VERSION) ./services/metrics --push

# Run services locally
run:
	@echo "Starting services..."
	docker compose up -d

# Stop services
stop:
	@echo "Stopping services..."
	docker compose down

# Development setup
dev-setup:
	@echo "Setting up development environment..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest

# Lint code
lint:
	@echo "Linting code..."
	@for service in auth gpio metrics; do \
		echo "Linting $$service..."; \
		cd services/$$service && golangci-lint run ./...; \
		cd $(GOBASE); \
	done

# Generate test coverage report
coverage:
	@echo "Generating coverage report..."
	@for service in auth gpio metrics; do \
		echo "Coverage for $$service..."; \
		cd services/$$service && go test -coverprofile=coverage.out ./... && \
		go tool cover -html=coverage.out -o coverage.html; \
		cd $(GOBASE); \
	done