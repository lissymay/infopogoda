# Variables
BINARY_NAME=weather
BIN_DIR=bin
CONFIG_PATH=config/config.yaml
CONFIG_REDIS_PATH=config/config-redis.yaml

# Default target
.PHONY: all
all: run

# ============================================
# RUN APPLICATION
# ============================================

.PHONY: run
run:
	go run ./cmd/linux/cli/main.go

.PHONY: run-redis
run-redis:
	go run ./cmd/linux/cli/main.go -config $(CONFIG_REDIS_PATH)

.PHONY: run-config
run-config:
	go run ./cmd/linux/cli/main.go -config $(CONFIG_PATH)

.PHONY: help
help:
	go run ./cmd/linux/cli/main.go -help

# ============================================
# BUILD
# ============================================

.PHONY: build
build:
	@echo "Building application..."
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY_NAME) ./cmd/linux/cli/main.go
	@echo "Build complete: $(BIN_DIR)/$(BINARY_NAME).exe"

# ============================================
# TESTS
# ============================================

.PHONY: test
test:
	@echo "Running all tests..."
	go test -v ./...

.PHONY: test-cover
test-cover:
	@echo "Running tests with coverage..."
	go test -cover ./...
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report saved to coverage.html"

.PHONY: test-cache
test-cache:
	@echo "Testing cache..."
	go test -v ./internal/pkg/cache/

.PHONY: test-config
test-config:
	@echo "Testing config..."
	go test -v ./pkg/config/

.PHONY: test-weather
test-weather:
	@echo "Testing weather adapter..."
	go test -v ./internal/adapters/weather/

.PHONY: test-cli
test-cli:
	@echo "Testing CLI..."
	go test -v ./internal/pkg/app/cli/

.PHONY: test-short
test-short:
	@echo "Running short tests..."
	go test -short -v ./...

# ============================================
# REDIS SERVICES
# ============================================

.PHONY: redis-start
redis-start:
	@echo "Starting Redis..."
	@if docker ps -a --format '{{.Names}}' | findstr /C:"redis-cache" > nul; then \
		echo "Container redis-cache already exists. Starting..."; \
		docker start redis-cache; \
	else \
		echo "Creating new container..."; \
		docker run -d --name redis-cache -p 6379:6379 redis:latest; \
	fi
	@echo "Redis started on port 6379"

.PHONY: redis-stop
redis-stop:
	@echo "Stopping Redis..."
	@docker stop redis-cache 2>nul || echo "Container not running"

.PHONY: redis-remove
redis-remove:
	@echo "Removing Redis container..."
	@docker stop redis-cache 2>nul
	@docker rm redis-cache 2>nul
	@echo "Container removed"

.PHONY: redis-status
redis-status:
	@docker ps -a | findstr redis-cache || echo "Container not found"

.PHONY: redis-clear
redis-clear:
	@echo "Clearing Redis cache..."
	@docker exec -it redis-cache redis-cli FLUSHALL
	@echo "Cache cleared"

# ============================================
# UTILITIES
# ============================================

.PHONY: clean
clean:
	@echo "Cleaning..."
	@if exist $(BIN_DIR) rmdir /s /q $(BIN_DIR)
	@if exist coverage.out del coverage.out
	@if exist coverage.html del coverage.html
	@echo "Clean complete"

.PHONY: deps
deps:
	@echo "Installing dependencies..."
	go mod tidy
	go mod download
	@echo "Dependencies installed"

.PHONY: fmt
fmt:
	@echo "Formatting code..."
	go fmt ./...

.PHONY: tidy
tidy:
	@echo "Tidying go.mod..."
	go mod tidy

# ============================================
# INFO
# ============================================

.PHONY: info
info:
	@echo "=========================================="
	@echo "Weather Info CLI Application"
	@echo "=========================================="
	@echo "Commands:"
	@echo "  make run           - Run with in-memory cache"
	@echo "  make run-redis     - Run with Redis cache"
	@echo "  make build         - Build binary"
	@echo "  make test          - Run all tests"
	@echo "  make test-cover    - Run tests with coverage"
	@echo "  make redis-start   - Start Redis (Docker)"
	@echo "  make redis-stop    - Stop Redis"
	@echo "  make clean         - Clean build artifacts"
	@echo "=========================================="