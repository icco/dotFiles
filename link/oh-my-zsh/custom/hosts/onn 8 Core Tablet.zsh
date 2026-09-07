export GOPATH=~/Projects/
export PATH="$GOPATH/bin:$PATH"
export GO111MODULE=on

# nvm without nvm_auto; see nvm_init_lazy in globals.zsh
nvm_init_lazy

# GPG
PINENTRY_USER_DATA="USE_CURSES=1"
GPG_TTY=$(tty)
export GPG_TTY

[ -f ~/.fzf.zsh ] && source ~/.fzf.zsh
# source /usr/share/doc/fzf/examples/key-bindings.zsh

# python
export PATH="$PATH:$HOME/.local/bin"

# zmv
autoload -Uz zmv

# vim: set filetype=zsh:
