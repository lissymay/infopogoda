<<<<<<< HEAD
# Variables
CONFIG_PATH=config/config.yaml
CONFIG_POGODA_PATH=config/config-pogoda.yaml

# Default target
.PHONY: all
all: run

# ============================================
# RUN APPLICATION
# ============================================

.PHONY: run
run:
	go run ./cmd/linux/cli/main.go

.PHONY: run-pogoda
run-pogoda:
	go run ./cmd/linux/cli/main.go -config $(CONFIG_POGODA_PATH)

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
	go build -o bin/weather ./cmd/linux/cli/main.go
	@echo "Build complete: bin/weather.exe"

.PHONY: build-pogoda
build-pogoda:
	@echo "Building application with pogoda.by..."
	go build -o bin/weather-pogoda ./cmd/linux/cli/main.go
	@echo "Build complete: bin/weather-pogoda.exe"

# ============================================
# UTILITIES
# ============================================

.PHONY: clean
clean:
	@echo "Cleaning..."
	@if exist bin rmdir /s /q bin
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
	@echo "  make run           - Run with open-meteo (default)"
	@echo "  make run-pogoda    - Run with pogoda.by"
	@echo "  make build         - Build binary"
	@echo "  make clean         - Clean build artifacts"
	@echo "  make deps          - Install dependencies"
	@echo "=========================================="
=======
.PHONY: run help build clean

run:
	go run ./cmd/linux/cli/main.go

run-config:
	go run ./cmd/linux/cli/main.go -config $(CONFIG)

help:
	go run ./cmd/linux/cli/main.go -help

build:
	go build -o bin/weather.exe ./cmd/linux/cli/main.go

clean:
	if exist bin rmdir /s /q bin
>>>>>>> aecab4d0c3cc7544596581e380fd0a9f06d16b88
