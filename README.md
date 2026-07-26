# Nat Welch's .Files

This repository holds my config files for just about everything. Use this to set up a fresh macOS or Linux box, and Nat will be a happy user.

Originally I did this entirely with shell scripts, then Ruby/Rake. Now I use Go for better performance and easier distribution.

## Prerequisites

- Go (see `go.mod` for the minimum version)
- [Task](https://taskfile.dev)
- git, with an SSH key registered on GitHub (`task vim` clones plugins over `git@github.com:`)
- Homebrew (macOS only, for `brew bundle`)

## Usage

```bash
# Install dotfiles
task infect

# Update vim plugins
task vim

# Update oh-my-zsh
task omz

# Install system packages (macOS)
brew bundle
```

For more options: `go run ./dotool --help` or `task --list`.

## Structure

- `link/` - Files symlinked to `~/.{filename}`
- `specific/` - Files symlinked preserving directory structure
- `bin/` - Scripts symlinked to `~/bin/`
