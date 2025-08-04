# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

### Development Commands
```bash
# Show usage help
go run ./dotool --help
make help

# Install dotfiles and link all configuration files
go run ./dotool infect
make infect

# Update vim plugins and sort spell file
go run ./dotool vim
make vim

# Run tests
go run ./dotool test
make test

# Build standalone binary
make build

# Install OSX packages (requires Homebrew)
make brew
```

## Architecture

This is a personal dotfiles repository that manages configuration files across multiple machines using Go. The system works by creating symbolic links from the repository to the user's home directory.

### Key Components

- **`dotool/`** - Go application that manages dotfile installation and vim plugin updates
  - `main.go` - Entry point using Cobra CLI framework with commands: infect, vim, test
  - `infect.go` - Core linking logic that creates directory structure and symlinks
  - `vim_impl.go` - Vim-specific functionality for plugin management and spell file sorting
  - `go.mod` - Module dependencies including github.com/spf13/cobra

- **`go.mod` (root)** - Main module file at repository root (github.com/icco/dotfiles)

- **`Makefile`** - Simple make targets that delegate to dotool commands without duplicating documentation

- **`link/`** - Files that get symlinked to `~/.{filename}` (with dot prefix added)
  - Contains shell configurations (bashrc, zshrc), vim config, git config, etc.
  - `oh-my-zsh/` - Complete Oh My Zsh installation with custom themes, plugins, and host-specific configurations

- **`specific/`** - Files that get symlinked with their directory structure preserved to `~/.{path}`
  - Used for nested config files like `.config/neofetch/config.conf`
  - Includes SSH configs, cron files, and pandoc filters

- **`bin/`** - Executable scripts that get symlinked to `~/bin/`

### Recent Changes

- **Cobra CLI Framework**: Replaced manual command parsing with professional `github.com/spf13/cobra` library
- **Simplified Makefile**: Removed documentation duplication - Makefile just calls dotool commands, Cobra handles help
- **Module Structure**: Root go.mod at repository level, separate go.mod in dotool/ with cobra dependency
- **Streamlined README**: Simplified documentation, removed outdated information

### Linking Behavior

The dotool infect command:
1. Creates backup copies of existing files in `~/tmp/` before overwriting
2. Creates absolute symlinks to repository files
3. Builds standard directory structure (`~/Projects`, `~/bin`, `~/tmp`)

### Vim Plugin Management

Vim plugins are managed by cloning fresh copies from GitHub, removing `.git` directories, and committing them directly to the dotfiles repo. This approach ensures reproducible plugin versions across machines without requiring external plugin managers.