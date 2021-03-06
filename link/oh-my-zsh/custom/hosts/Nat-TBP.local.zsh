# ZSH config for work mbp

export TZ="UTC"

# Add iterm2 support
test -e "${HOME}/.iterm2_shell_integration.zsh" && source "${HOME}/.iterm2_shell_integration.zsh"

# Try and fix PATH
export PATH="/usr/local/sbin:/usr/local/bin:$PATH"

# Mac ls does not have color option...
alias ls="/bin/ls";

# for https://github.com/firstlookmedia/aws-profile-gpg
export AWS_PROFILE_GPG_HOME=$HOME/Projects/aws-profile-gpg

[ -f /usr/local/etc/profile.d/z.sh ] && source /usr/local/etc/profile.d/z.sh
export PATH="/usr/local/opt/gettext/bin:$PATH"

# Colors
export CLICOLOR=1
export LSCOLORS=ExFxCxDxBxegedabagacad

# Golang
export GOPATH="$HOME/Projects"
export PATH="$GOPATH/bin:$PATH"
export GO111MODULE="on"

# RVM
[[ -s "$HOME/.rvm/scripts/rvm" ]] && source "$HOME/.rvm/scripts/rvm" # Load RVM into a shell session *as a function*
export PATH="$PATH:$HOME/.rvm/bin"

# Add GVM
#[[ -s "$HOME/.gvm/scripts/gvm" ]] && source "${HOME}/.gvm/scripts/gvm"

# added by travis gem
[ -f /Users/nat/.travis/travis.sh ] && source /Users/nat/.travis/travis.sh

# For jumping
[ -f /usr/local/etc/profile.d/z.sh ] && . /usr/local/etc/profile.d/z.sh

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

# The next line updates PATH for the Google Cloud SDK.
if [ -f '/Users/nat/google-cloud-sdk/path.zsh.inc' ]; then source '/Users/nat/google-cloud-sdk/path.zsh.inc'; fi

# The next line enables shell command completion for gcloud.
if [ -f '/Users/nat/google-cloud-sdk/completion.zsh.inc' ]; then source '/Users/nat/google-cloud-sdk/completion.zsh.inc'; fi

[ -f ~/.fzf.zsh ] && source ~/.fzf.zsh

alias tf="terraform"

# vim: set filetype=zsh:
