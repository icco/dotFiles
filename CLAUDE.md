# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

### Development Commands
```bash
# Install dotfiles and link all configuration files
go run ./infect

# Update vim plugins and sort spell file
go run ./infect vim

# Run tests
go run ./infect test

# Install OSX packages (requires Homebrew)
brew bundle
```

## Architecture

This is a personal dotfiles repository that manages configuration files across multiple machines using Go. The system works by creating symbolic links from the repository to the user's home directory.

### Key Components

- **`infect/`** - Go application that manages dotfile installation and vim plugin updates
  - `main.go` - Entry point with commands: infect (default), vim, test
  - `infect.go` - Core linking logic that creates directory structure and symlinks
  - `vim_impl.go` - Vim-specific functionality for plugin management and spell file sorting

- **`link/`** - Files that get symlinked to `~/.{filename}` (with dot prefix added)
  - Contains shell configurations (bashrc, zshrc), vim config, git config, etc.
  - `oh-my-zsh/` - Complete Oh My Zsh installation with custom themes and plugins

- **`specific/`** - Files that get symlinked with their directory structure preserved to `~/.{path}`
  - Used for nested config files like `.config/neofetch/config.conf`

- **`bin/`** - Executable scripts that get symlinked to `~/bin/`

### Linking Behavior

The infect tool:
1. Creates backup copies of existing files in `~/tmp/` before overwriting
2. Creates absolute symlinks to repository files
3. Builds standard directory structure (`~/Projects`, `~/bin`, `~/tmp`)

### Vim Plugin Management

Vim plugins are managed by cloning fresh copies from GitHub, removing `.git` directories, and committing them directly to the dotfiles repo. This approach ensures reproducible plugin versions across machines without requiring external plugin managers.