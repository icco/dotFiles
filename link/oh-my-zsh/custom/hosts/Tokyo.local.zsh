# zshrc for my MBP

export TZ="UTC"

# Add iterm2 support
test -e "${HOME}/.iterm2_shell_integration.zsh" && source "${HOME}/.iterm2_shell_integration.zsh"

# Try and fix PATH
export PATH="/usr/local/sbin:/usr/local/bin:$PATH"

# Mac ls does not have color option...
alias ls="/bin/ls";

# Colors
export CLICOLOR=1
export LSCOLORS=ExFxCxDxBxegedabagacad

# Golang
export GOPATH="$HOME/Projects"
export PATH="$GOPATH/bin:$PATH"

# Add RVM to PATH for scripting. Make sure this is the last PATH variable change.
export PATH="$PATH:$HOME/.rvm/bin"
[[ -s "$HOME/.rvm/scripts/rvm" ]] && source "$HOME/.rvm/scripts/rvm" # Load RVM into a shell session *as a function*

# Add GVM
#[[ -s "$HOME/.gvm/scripts/gvm" ]] && source "${HOME}/.gvm/scripts/gvm"

# added by travis gem
[ -f /Users/nat/.travis/travis.sh ] && source /Users/nat/.travis/travis.sh

# For jumping
[ -f /usr/local/etc/profile.d/z.sh ] && . /usr/local/etc/profile.d/z.sh

## For yubikey life
#function init_gpg_ssh {
#  source ~/.gpg-agent-info;
#  for key in $( cat ~/.gpg-agent-info | cut -d = -f 1 ); do
#    eval "export $key"
#  done
#  ssh-add -l 2> /dev/null
#}
#init_gpg_ssh

# A better which
alias which='alias | /usr/local/bin/gwhich --tty-only --read-alias --show-dot --show-tilde'

export NVM_DIR="$HOME/.nvm"
. "/usr/local/opt/nvm/nvm.sh"

# place this after nvm initialization!
autoload -U add-zsh-hook
load-nvmrc() {
  local node_version="$(nvm version)"
  local nvmrc_path="$(nvm_find_nvmrc)"

  if [ -n "$nvmrc_path" ]; then
    local nvmrc_node_version=$(nvm version "$(cat "${nvmrc_path}")")

    if [ "$nvmrc_node_version" = "N/A" ]; then
      nvm install
    elif [ "$nvmrc_node_version" != "$node_version" ]; then
      nvm use
    fi
  elif [ "$node_version" != "$(nvm version default)" ]; then
    echo "Reverting to nvm default version"
    nvm use default
  fi
}
add-zsh-hook chpwd load-nvmrc
load-nvmrc

# vim: set filetype=zsh:
