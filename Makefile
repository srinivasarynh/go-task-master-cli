.PHONY: build install test clean lint run

APP_NAME=taskmaster
BUILD_DIR=build
INSTALL_DIR=/usr/local/bin

# Build the application
build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) .

# Install the application
install: build
	@echo "Installing $(APP_NAME) to $(INSTALL_DIR)..."
	@sudo cp $(BUILD_DIR)/$(APP_NAME) $(INSTALL_DIR)/

# Run tests
test:
	@echo "Running tests..."
	go test ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -cover ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)

# Lint code
lint:
	@echo "Running linter..."
	golangci-lint run

# Run the application
run:
	go run main.go

# Build for multiple platforms
build-all:
	@echo "Building for multiple platforms..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-linux-amd64 .
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-darwin-amd64 .
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-windows-amd64.exe .

# Create release
release: clean build-all
	@echo "Creating release..."
	@cd $(BUILD_DIR) && tar -czf $(APP_NAME)-linux-amd64.tar.gz $(APP_NAME)-linux-amd64
	@cd $(BUILD_DIR) && tar -czf $(APP_NAME)-darwin-amd64.tar.gz $(APP_NAME)-darwin-amd64
	@cd $(BUILD_DIR) && zip $(APP_NAME)-windows-amd64.zip $(APP_NAME)-windows-amd64.exe

# Development mode
dev:
	@echo "Starting development mode..."
	air
