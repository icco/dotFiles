# ZSH config for work mbp
export TZ="UTC"

# Add iterm2 support
test -e "${HOME}/.iterm2_shell_integration.zsh" && source "${HOME}/.iterm2_shell_integration.zsh"

# Try and fix PATH
export PATH="/opt/homebrew/bin:/usr/local/sbin:/usr/local/bin:$PATH"

# Mac ls does not have color option...
alias ls="/bin/ls";

export PATH="/usr/local/opt/gettext/bin:$PATH"

# Colors
export CLICOLOR=1
export LSCOLORS=ExFxCxDxBxegedabagacad

# Golang
export GOPATH="$HOME/Projects"
export PATH="$GOPATH/bin:$PATH"
export GO111MODULE="on"
export GOPRIVATE="github.com/pinginc/*"

# search!
[ -f ~/.fzf.zsh ] && source ~/.fzf.zsh
alias fr='open -R "$(fzf)"'
alias f='open "$(fzf)"'
alias fv='vim "$(fzf)"'

alias k='kubectl'

wordcount() {
  pandoc --lua-filter wordcount.lua "$@"
}

# gallery + ytdl
alias gd=gallery-dl

# A better which
alias which='alias | gwhich --tty-only --read-alias --show-dot --show-tilde'

# iaWriter
alias ia='open -a "IA Writer"'

# nvm
export NVM_DIR="$HOME/.nvm"
[ -s "/opt/homebrew/opt/nvm/nvm.sh" ] && \. "/opt/homebrew/opt/nvm/nvm.sh"  # This loads nvm
[ -s "/opt/homebrew/opt/nvm/etc/bash_completion.d/nvm" ] && \. "/opt/homebrew/opt/nvm/etc/bash_completion.d/nvm"  # This loads nvm bash_completion

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

# Terraform
alias tf="tofu"
alias tfdocs="terraform-docs"

# Github
alias ghr='gh repo view --web'
alias ghpr='gh pr view --web'
alias ghpc='gh pr create --web'

# RVM
[[ -s "$HOME/.rvm/scripts/rvm" ]] && source "$HOME/.rvm/scripts/rvm" # Load RVM into a shell session *as a function*
export PATH="$PATH:$HOME/.rvm/bin"

# zmv
autoload -Uz zmv

# vim: set filetype=zsh:
