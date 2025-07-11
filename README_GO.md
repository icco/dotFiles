# Dotfiles Manager - Go Edition

This repository has been migrated from Ruby/Rake to Go for better performance, easier distribution, and improved maintainability.

## Migration Overview

The original Ruby-based dotfiles manager used:
- `Rakefile` - Main task definitions
- `Gemfile` - Ruby dependencies
- `rake` commands for operations

The new Go-based dotfiles manager provides:
- `main.go` - Main application entry point
- `infect.go` - File linking and directory structure management
- `vim.go` - Vim spell sorting and plugin management
- Comprehensive test suite
- Easy build and distribution

## Commands

### Original Ruby Commands → New Go Commands

| Ruby Command | Go Command | Description |
|--------------|------------|-------------|
| `rake` | `./dotfiles infect` | Hook dotfiles into system-standard positions |
| `rake vim` | `./dotfiles vim` | Sort vim spell and upgrade vim plugins |
| `rake test` | `./dotfiles test` | Test to make sure everything works ok |

## Quick Start

1. **Build the executable:**
   ```bash
   make build
   # or
   go build -o dotfiles .
   ```

2. **Run the main command (equivalent to `rake`):**
   ```bash
   ./dotfiles infect
   ```

3. **Run vim command (equivalent to `rake vim`):**
   ```bash
   ./dotfiles vim
   ```

## Using Make

The included `Makefile` provides convenient shortcuts:

```bash
# Build and run infect command
make run-infect

# Build and run vim command  
make run-vim

# Run tests
make test

# Install to system PATH
make install

# Show all available commands
make help
```

## Installation

### Option 1: Local Build
```bash
make build
./dotfiles <command>
```

### Option 2: System Installation
```bash
make install
dotfiles <command>
```

## Testing

Run the comprehensive test suite:

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage
```

The test suite includes:
- File linking functionality
- Directory structure creation
- Vim spell sorting
- Plugin management (with proper error handling)
- Symlink creation and backup

## Features

### Infect Command (`./dotfiles infect`)
- Creates required directory structure (`~/Projects`, `~/bin`, `~/tmp`)
- Links all files from `link/` directory to home directory with dot prefix
- Links all files from `specific/` directory maintaining directory structure
- Links all files from `bin/` directory to `~/bin`
- Backs up existing files before overwriting
- Creates symbolic links with absolute paths

### Vim Command (`./dotfiles vim`)
- Sorts and deduplicates vim spell file (`link/vim/spell/en.utf-8.add`)
- Commits spell changes to git
- Upgrades all vim plugins by cloning fresh copies
- Removes `.git` directories from plugins
- Commits all plugin changes to git

### Test Command (`./dotfiles test`)
- Performs basic validation checks
- Ensures Go version compatibility
- Validates environment setup

## Error Handling

The Go version includes improved error handling:
- Graceful handling of missing files/directories
- Proper backup creation before overwriting
- Detailed error messages with context
- Safe fallbacks for common failure scenarios

## Benefits of Go Migration

1. **Performance**: Faster execution compared to Ruby
2. **Distribution**: Single binary, no runtime dependencies
3. **Cross-platform**: Easy compilation for different operating systems
4. **Testing**: Comprehensive test suite with good coverage
5. **Maintainability**: Type-safe, compiled language with better tooling
6. **Dependencies**: No external dependencies beyond Go standard library

## Migration Notes

- All original functionality has been preserved
- Command-line interface remains similar for easy transition
- Git operations and file system operations work identically
- Backup functionality is enhanced with timestamped backups
- Error messages are more detailed and helpful

## Troubleshooting

### Common Issues

1. **Permission denied errors:**
   - Ensure you have write permissions to your home directory
   - Use `sudo` if installing system-wide

2. **Git authentication errors:**
   - Ensure SSH keys are properly configured for GitHub
   - Plugin upgrades will fail gracefully if SSH is not configured

3. **Symlink creation fails:**
   - On Windows, ensure Developer Mode is enabled
   - On macOS, ensure you have necessary permissions

### Getting Help

```bash
# Show usage information
./dotfiles

# Run tests to verify setup
make test

# Check for common issues
./dotfiles test
``` 