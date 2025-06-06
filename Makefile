# Wu-Tang Countdown Timer Makefile
# Wu-Tang Clan ain't nuthin' ta f' wit!

# Variables
BINARY_NAME=wut
MAIN_FILE=cmd/wut/main.go
BUILD_DIR=bin
RELEASE_DIR=build
VERSION?=1.0.0

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Build flags
LDFLAGS=-ldflags="-s -w -X main.version=$(VERSION)"

# Git info for version
GIT_COMMIT=$(shell git rev-parse --short HEAD 2>/dev/null || echo "dev")
GIT_TAG=$(shell git describe --tags --exact-match 2>/dev/null || echo "")
ifneq ($(GIT_TAG),)
	VERSION=$(GIT_TAG)
else
	VERSION=$(GIT_COMMIT)
endif

# Default target
.PHONY: all
all: clean deps build

# Initialize and install dependencies
.PHONY: deps
deps:
	@echo "🔥 Installing Wu-Tang dependencies..."
	$(GOMOD) init wu-tang-timer 2>/dev/null || true
	$(GOGET) github.com/charmbracelet/bubbletea
	$(GOGET) github.com/charmbracelet/lipgloss
	$(GOGET) github.com/charmbracelet/bubbles/progress
	$(GOMOD) tidy
	@echo "✅ Dependencies installed!"

# Build the binary
.PHONY: build
build:
	@echo "🐉 Building Wu-Tang Timer..."
	mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/wut
	@echo "✅ Binary built: $(BUILD_DIR)/$(BINARY_NAME)"

# Build optimized release version
.PHONY: release
release:
	@echo "🏆 Building Wu-Tang Timer release..."
	mkdir -p $(RELEASE_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(RELEASE_DIR)/$(BINARY_NAME) ./cmd/wut
	@echo "✅ Release binary built: $(RELEASE_DIR)/$(BINARY_NAME)"

# Cross-platform builds
.PHONY: build-all
build-all: clean
	@echo "🌍 Building Wu-Tang Timer for all platforms..."
	mkdir -p $(RELEASE_DIR)

	# macOS ARM64 (M1/M2)
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o $(RELEASE_DIR)/$(BINARY_NAME)-macos-m1 ./cmd/wut

	# macOS Intel
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(RELEASE_DIR)/$(BINARY_NAME)-macos-intel ./cmd/wut

	# Linux
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(RELEASE_DIR)/$(BINARY_NAME)-linux ./cmd/wut

	# Windows
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(RELEASE_DIR)/$(BINARY_NAME).exe ./cmd/wut

	@echo "✅ All platform binaries built in $(RELEASE_DIR)/"
	@ls -la $(RELEASE_DIR)/

# Run the application
.PHONY: run
run: deps
	@echo "🚀 Running Wu-Tang Timer..."
	$(GOCMD) run ./cmd/wut

# Run with hot reload (requires air)
.PHONY: dev
dev:
	@if command -v air > /dev/null; then \
		echo "🔄 Running with hot reload..."; \
		air; \
	else \
		echo "⚠️  Air not installed. Install with: go install github.com/cosmtrek/air@latest"; \
		echo "🚀 Running normally..."; \
		$(MAKE) run; \
	fi

# Test the application
.PHONY: test
test:
	@echo "🧪 Testing Wu-Tang Timer..."
	$(GOTEST) -v ./...

# Clean build artifacts
.PHONY: clean
clean:
	@echo "🧹 Cleaning build artifacts..."
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)
	rm -rf $(RELEASE_DIR)
	@echo "✅ Clean complete!"

# Install binary to system PATH
.PHONY: install
install: build
	@echo "📦 Installing Wu-Tang Timer to system..."
	cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/wut
	@echo "✅ Installed! Run with: wut"

# Uninstall binary from system PATH
.PHONY: uninstall
uninstall:
	@echo "🗑️  Uninstalling Wu-Tang Timer..."
	rm -f /usr/local/bin/wut
	@echo "✅ Uninstalled!"

# Format code
.PHONY: fmt
fmt:
	@echo "🎨 Formatting code..."
	$(GOCMD) fmt ./...

# Lint code (requires golangci-lint)
.PHONY: lint
lint:
	@if command -v golangci-lint > /dev/null; then \
		echo "🔍 Linting code..."; \
		golangci-lint run; \
	else \
		echo "⚠️  golangci-lint not installed. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

# Show binary info
.PHONY: info
info:
	@echo "📊 Wu-Tang Timer Info:"
	@echo "  Binary: $(BINARY_NAME)"
	@echo "  Version: $(VERSION)"
	@echo "  Main file: $(MAIN_FILE)"
	@echo "  Build dir: $(BUILD_DIR)"
	@if [ -f $(BUILD_DIR)/$(BINARY_NAME) ]; then \
		echo "  Size: $(ls -lah $(BUILD_DIR)/$(BINARY_NAME) | awk '{print $5}')"; \
		echo "  Built: $(stat -f '%Sm' $(BUILD_DIR)/$(BINARY_NAME) 2>/dev/null || stat -c '%y' $(BUILD_DIR)/$(BINARY_NAME) 2>/dev/null)"; \
	else \
		echo "  Status: Not built"; \
	fi

# Show help
.PHONY: help
help:
	@echo "🐉 Wu-Tang Countdown Timer Makefile"
	@echo ""
	@echo "Usage:"
	@echo "  make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  all        - Clean, install deps, and build"
	@echo "  deps       - Initialize module and install dependencies"
	@echo "  build      - Build binary for current platform"
	@echo "  release    - Build optimized release binary"
	@echo "  build-all  - Build for all platforms (macOS, Linux, Windows)"
	@echo "  run        - Install deps and run the application"
	@echo "  dev        - Run with hot reload (requires air)"
	@echo "  test       - Run tests"
	@echo "  clean      - Remove build artifacts"
	@echo "  install    - Install binary to /usr/local/bin"
	@echo "  uninstall  - Remove binary from /usr/local/bin"
	@echo "  fmt        - Format Go code"
	@echo "  lint       - Lint code (requires golangci-lint)"
	@echo "  info       - Show binary information"
	@echo "  help       - Show this help message"
	@echo ""
	@echo "Wu-Tang Clan ain't nuthin' ta f' wit! 🔥"
