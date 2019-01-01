[[ -s "$HOME/.rvm/scripts/rvm" ]] && source "$HOME/.rvm/scripts/rvm" # Load RVM into a shell session *as a function*

[ -f ~/.fzf.bash ] && source ~/.fzf.bash

# Golang
export GOPATH=~/Projects/
export PATH="~/Projects/bin:$PATH"

# GPG
PINENTRY_USER_DATA="USE_CURSES=1"
GPG_TTY=$(tty)
export GPG_TTY

. /home/nat/Projects/z/z.sh
