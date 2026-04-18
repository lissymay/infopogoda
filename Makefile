# Variables
CONFIG_PATH=config/config.yaml
CONFIG_POGODA_PATH=config/config-pogoda.yaml

# Default target
.PHONY: all
all: run

# ============================================
# RUN CLI APPLICATION
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
# RUN GUI APPLICATION
# ============================================

.PHONY: run-gui
run-gui:
	go run ./cmd/linux/gui/main.go

.PHONY: run-gui-pogoda
run-gui-pogoda:
	go run ./cmd/linux/gui/main.go -config $(CONFIG_POGODA_PATH)

.PHONY: run-gui-config
run-gui-config:
	go run ./cmd/linux/gui/main.go -config $(CONFIG_PATH)

# ============================================
# BUILD CLI
# ============================================

.PHONY: build
build:
	@echo "Building CLI application..."
	go build -o bin/weather-cli ./cmd/linux/cli/main.go
	@echo "Build complete: bin/weather-cli.exe"

.PHONY: build-pogoda
build-pogoda:
	@echo "Building CLI application with pogoda.by..."
	go build -o bin/weather-cli-pogoda ./cmd/linux/cli/main.go
	@echo "Build complete: bin/weather-cli-pogoda.exe"

# ============================================
# BUILD GUI
# ============================================

.PHONY: build-gui
build-gui:
	@echo "Building GUI application..."
	go build -o bin/weather-gui ./cmd/linux/gui/main.go
	@echo "Build complete: bin/weather-gui.exe"

.PHONY: build-gui-pogoda
build-gui-pogoda:
	@echo "Building GUI application with pogoda.by..."
	go build -o bin/weather-gui-pogoda ./cmd/linux/gui/main.go
	@echo "Build complete: bin/weather-gui-pogoda.exe"

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
	@echo "Weather Info Application"
	@echo "=========================================="
	@echo ""
	@echo "CLI Commands:"
	@echo "  make run              - Run CLI with open-meteo"
	@echo "  make run-pogoda       - Run CLI with pogoda.by"
	@echo "  make build            - Build CLI binary"
	@echo ""
	@echo "GUI Commands:"
	@echo "  make run-gui          - Run GUI with open-meteo"
	@echo "  make run-gui-pogoda   - Run GUI with pogoda.by"
	@echo "  make build-gui        - Build GUI binary"
	@echo ""
	@echo "Utilities:"
	@echo "  make clean            - Clean build artifacts"
	@echo "  make deps             - Install dependencies"
	@echo "  make fmt              - Format code"
	@echo "=========================================="