[[ -s "$HOME/.rvm/scripts/rvm" ]] && source "$HOME/.rvm/scripts/rvm" # Load RVM into a shell session *as a function*

# Golang
export GOPATH=~/Projects/
export PATH="$GOPATH/bin:$PATH"
export GO111MODULE=on

export PATH="/snap/bin:$PATH"

export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"

# GPG
PINENTRY_USER_DATA="USE_CURSES=1"
GPG_TTY=$(tty)
export GPG_TTY

[ -f ~/.fzf.zsh ] && source ~/.fzf.zsh
source /usr/share/doc/fzf/examples/key-bindings.zsh

. /home/nat/Projects/z/z.sh
