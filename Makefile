.PHONY: build clean run

# Binary name and directories
BINARY=slack_jack
BIN_DIR=bin
CMD_DIR=cmd/slack_jack

# Create bin directory if it doesn't exist
$(BIN_DIR):
	mkdir -p $(BIN_DIR)

# Build the binary
build: $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY) ./$(CMD_DIR)

# Clean build artifacts
clean:
	go clean
	rm -rf $(BIN_DIR)

# Run the application (requires token)
run: build
	./$(BIN_DIR)/$(BINARY) -t $(token)

# Install dependencies
deps:
	go mod download

# Build for multiple platforms
build-all: clean $(BIN_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/$(BINARY)_linux_amd64 ./$(CMD_DIR)
	GOOS=darwin GOARCH=amd64 go build -o $(BIN_DIR)/$(BINARY)_darwin_amd64 ./$(CMD_DIR)
	GOOS=windows GOARCH=amd64 go build -o $(BIN_DIR)/$(BINARY)_windows_amd64.exe ./$(CMD_DIR)

# Default target
all: clean build
