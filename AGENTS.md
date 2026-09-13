# AGENTS.md

Guidance for coding agents working on dotFiles.

## Project Overview

Personal configuration and dotfiles repository managed by a custom Go CLI utility `dotool` (`github.com/icco/dotfiles`).

## Commands (Taskfile)

Run via `task <name>`:
- `task build` — Build `dotool` binary to `bin/dotool`
- `task test` — Run Go tests (`go test ./...`)
- `task infect` — Install / symlink dotfiles into home directory
- `task vim` — Update vim plugins and sort spellfile
- `task omz` — Update Oh My Zsh
- `task brew` — Install macOS packages via Homebrew bundle
- `task clean` — Clean built binaries

## Architecture & Layout

- `dotool/` — Go CLI source code for dotfile management.
- `link/` — Files symlinked into `$HOME`.
- `copy/` — Files copied into `$HOME`.

## Conventions

- PR titles and commits must follow Conventional Commits with lowercase subjects.
- Ensure `go test ./...` passes for any changes to `dotool`.
- Never commit private keys, tokens, or plaintext credentials.
