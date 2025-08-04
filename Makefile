.PHONY: help infect vim test omz build clean brew

# Default target - let dotool handle help
help:
	@go run ./dotool --help

# Install dotfiles
infect:
	@go run ./dotool infect

# Update vim plugins and sort spell file
vim:
	@go run ./dotool vim

# Run tests
test:
	@go run ./dotool test

# Update oh-my-zsh
omz:
	@go run ./dotool omz

# Build the dotool binary
build:
	@go build -o bin/dotool ./dotool

# Clean built binaries
clean:
	@rm -f bin/dotool

# Install OSX packages using Homebrew
brew:
	@brew bundle