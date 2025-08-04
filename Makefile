.PHONY: help install vim test clean build brew

# Default target
help:
	@echo "dotool - Dotfiles management tool"
	@echo ""
	@echo "Available targets:"
	@echo "  help     Show this help message"
	@echo "  install  Install dotfiles and link all configuration files"
	@echo "  vim      Update vim plugins and sort spell file"
	@echo "  test     Run tests"
	@echo "  build    Build the dotool binary"
	@echo "  clean    Remove built binaries"
	@echo "  brew     Install OSX packages using Homebrew"
	@echo ""
	@echo "Examples:"
	@echo "  make install"
	@echo "  make vim"
	@echo "  make test"

# Install dotfiles
install:
	go run ./dotool install

# Update vim plugins and sort spell file
vim:
	go run ./dotool vim

# Run tests
test:
	go run ./dotool test

# Build the dotool binary
build:
	go build -o bin/dotool ./dotool

# Clean built binaries
clean:
	rm -f bin/dotool

# Install OSX packages using Homebrew
brew:
	brew bundle