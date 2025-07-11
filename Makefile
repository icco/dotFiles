.PHONY: build test clean run-infect run-vim run-test help

# Build the dotfiles executable
build:
	go build -o dotfiles .

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Clean build artifacts
clean:
	rm -f dotfiles coverage.out coverage.html

# Run the infect command
run-infect: build
	./dotfiles infect

# Run the vim command
run-vim: build
	./dotfiles vim

# Run the test command
run-test: build
	./dotfiles test

# Install the executable to a location in PATH
install: build
	cp dotfiles /usr/local/bin/dotfiles

# Uninstall the executable
uninstall:
	rm -f /usr/local/bin/dotfiles

# Show help
help:
	@echo "Available commands:"
	@echo "  build        - Build the dotfiles executable"
	@echo "  test         - Run tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  clean        - Remove build artifacts"
	@echo "  run-infect   - Build and run the infect command"
	@echo "  run-vim      - Build and run the vim command"
	@echo "  run-test     - Build and run the test command"
	@echo "  install      - Install executable to /usr/local/bin"
	@echo "  uninstall    - Remove executable from /usr/local/bin"
	@echo "  help         - Show this help message" 