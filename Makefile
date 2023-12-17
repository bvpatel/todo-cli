.PHONY: test build run clean

# Go related variables
GO_CMD := go
GO_TEST := $(GO_CMD) test
GO_BUILD := $(GO_CMD) build

# App variables
APP_NAME=todo-cli
MAIN_FILE=main.go

# Test the code
test:
	$(GO_TEST) -v ./...

# Build the application
build:
	$(GO_BUILD) -o $(APP_NAME) $(MAIN_FILE)

# Run the application
run:
	./$(APP_NAME)

# Clean build artifacts
clean:
	rm -f $(APP_NAME)
