# Nat Welch's .Files

This repository holds my config files for just about everything. Use this to set up a fresh OSX or Linux box, and Nat will be a happy user.

Originally I did this entirely with shell scripts, then Ruby/Rake. Now I use Go for better performance and easier distribution.

## Usage

```bash
# Install dotfiles
make infect

# Update vim plugins
make vim

# Install system packages
make brew
```

For more options: `go run ./dotool --help`

## Structure

- `link/` - Files symlinked to `~/.{filename}`
- `specific/` - Files symlinked preserving directory structure
- `bin/` - Scripts symlinked to `~/bin/`
