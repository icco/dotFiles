# tirith plugin for Oh-My-Zsh

Integrates [tirith](https://github.com/sheeki03/tirith) terminal security into your shell.

## Installation

1. Install tirith: `brew install sheeki03/tap/tirith`
2. Clone this repo:
   ```sh
   git clone https://github.com/sheeki03/ohmyzsh-tirith \
     ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/tirith
   ```
3. Add `tirith` to plugins in `~/.zshrc`:
   ```sh
   plugins=(... tirith)
   ```
4. Restart your terminal

## Configuration

By default, the plugin warns if tirith is not installed but allows shell startup to continue.
To enforce tirith presence (block shell startup if missing):

```sh
export TIRITH_STRICT=1
```

Add this before the `plugins=()` line in your `.zshrc`.
