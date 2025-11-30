# GoNeSh Makefile

# Variables
BINARY_NAME=gonesh
VERSION=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_DIR=build
GO_FILES=$(shell find . -name '*.go' -not -path './vendor/*')

# Go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=gofmt
GOLINT=golangci-lint

# Build flags
LDFLAGS=-ldflags "-X main.version=$(VERSION)"

.PHONY: all build run test lint fmt clean install help

# Default target
all: build

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/gonesh/

# Run the application
run: build
	@./$(BUILD_DIR)/$(BINARY_NAME)

# Run in development mode (rebuild on change would need external tool)
dev: build
	@./$(BUILD_DIR)/$(BINARY_NAME)

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

# Run linter
lint:
	@echo "Running linter..."
	$(GOLINT) run ./...

# Format code
fmt:
	@echo "Formatting code..."
	$(GOFMT) -s -w $(GO_FILES)

# Tidy dependencies
tidy:
	@echo "Tidying dependencies..."
	$(GOMOD) tidy

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html
	@rm -f $(BINARY_NAME)

# Install to GOPATH/bin
install: build
	@echo "Installing $(BINARY_NAME)..."
	@cp $(BUILD_DIR)/$(BINARY_NAME) $(GOPATH)/bin/

# Install to /usr/local/bin (requires sudo)
install-global: build
	@echo "Installing $(BINARY_NAME) to /usr/local/bin..."
	@sudo cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	$(GOMOD) download

# Update dependencies
deps-update:
	@echo "Updating dependencies..."
	$(GOGET) -u ./...
	$(GOMOD) tidy

# Show help
help:
	@echo "GoNeSh Makefile commands:"
	@echo ""
	@echo "  make build          - Build the binary"
	@echo "  make run            - Build and run"
	@echo "  make dev            - Run in development mode"
	@echo "  make test           - Run tests"
	@echo "  make test-coverage  - Run tests with coverage report"
	@echo "  make lint           - Run linter"
	@echo "  make fmt            - Format code"
	@echo "  make tidy           - Tidy go.mod"
	@echo "  make clean          - Clean build artifacts"
	@echo "  make install        - Install to GOPATH/bin"
	@echo "  make install-global - Install to /usr/local/bin"
	@echo "  make deps           - Download dependencies"
	@echo "  make deps-update    - Update dependencies"
	@echo "  make help           - Show this help"
