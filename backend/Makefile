build:
	@echo "Building the backend..."
	@go build -o bin/backend
	@echo "Build completed."

run: build
	@echo "Running the backend..."
	@./bin/backend

test:
	@echo "Running tests..."
	@go test -v ./...
	@echo "Tests completed."

install:
	@echo "Installing dependencies..."
	go install github.com/githubnemo/CompileDaemon@latest
	@echo "Dependencies installed."

dev:
	@echo "Starting server with auto-restart using CompileDaemon..."
	CompileDaemon -command="make run"

clean:
	@echo "Cleaning up..."
	go clean
	@echo "Cleanup completed."