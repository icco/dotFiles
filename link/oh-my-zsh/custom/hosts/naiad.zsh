
# Golang
export GOPATH=~/Projects/
export PATH="$GOPATH/bin:$PATH"

# GPG
PINENTRY_USER_DATA="USE_CURSES=1"
GPG_TTY=$(tty)
export GPG_TTY

# python
export PATH="$PATH:$HOME/.local/bin"

# zmv
autoload -Uz zmv

# vim: set filetype=zsh:
