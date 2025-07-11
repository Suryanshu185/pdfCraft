# Makefile for pdfCraft API

.PHONY: build test run clean docker-build docker-run help

# Build the application
build:
	go build -o pdfcraft-api cmd/api/main.go

# Run tests
test:
	go test -v ./...

# Run the application locally
run:
	go run cmd/api/main.go

# Clean build artifacts
clean:
	rm -f pdfcraft-api
	rm -rf uploads/ temp/ cache/ test_files/

# Format code
fmt:
	go fmt ./...

# Run linter
lint:
	golangci-lint run

# Install dependencies
deps:
	go mod download
	go mod tidy

# Build Docker image
docker-build:
	docker build -t pdfcraft-api .

# Run Docker container
docker-run:
	docker run -p 8080:8080 pdfcraft-api

# Run with Docker Compose
docker-compose-up:
	docker-compose up --build

# Run with Docker Compose in background
docker-compose-up-detached:
	docker-compose up -d --build

# Stop Docker Compose
docker-compose-down:
	docker-compose down

# Test API endpoints
test-api:
	./test_api.sh

# Help
help:
	@echo "Available targets:"
	@echo "  build                    - Build the application"
	@echo "  test                     - Run tests"
	@echo "  run                      - Run the application locally"
	@echo "  clean                    - Clean build artifacts"
	@echo "  fmt                      - Format code"
	@echo "  lint                     - Run linter"
	@echo "  deps                     - Install dependencies"
	@echo "  docker-build             - Build Docker image"
	@echo "  docker-run               - Run Docker container"
	@echo "  docker-compose-up        - Run with Docker Compose"
	@echo "  docker-compose-up-detached - Run with Docker Compose in background"
	@echo "  docker-compose-down      - Stop Docker Compose"
	@echo "  test-api                 - Test API endpoints"
	@echo "  help                     - Show this help"