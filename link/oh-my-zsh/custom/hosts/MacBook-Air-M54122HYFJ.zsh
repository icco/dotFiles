# ZSH config for work mbp
export TZ="UTC"

export ART_API_URL=https://art.natwelch.com
export ART_API_AUDIENCE=32555940559.apps.googleusercontent.com

# Add iterm2 support
test -e "${HOME}/.iterm2_shell_integration.zsh" && source "${HOME}/.iterm2_shell_integration.zsh"

# Try and fix PATH
export PATH="/opt/homebrew/bin:/usr/local/sbin:/usr/local/bin:$PATH"

# alias ls to GNU ls with colors
alias ls="/opt/homebrew/bin/gls --color";

# Colors
export CLICOLOR=1
export LSCOLORS=ExFxCxDxBxegedabagacad
export LS_COLORS=$(vivid generate iceberg-dark)

# Golang
export GOPATH="$HOME/Projects"
export PATH="$GOPATH/bin:$PATH"
export GO111MODULE="on"
export GOPRIVATE="github.com/pinginc/*"

# search! (fzf >= 0.48 integration: CTRL-R history, CTRL-T files, ALT-C cd, **<TAB> completion)
(( $+commands[fzf] )) && source <(fzf --zsh)
alias fr='open -R "$(fzf)"'
alias f='open "$(fzf)"'
alias fv='vim "$(fzf)"'

alias k='kubectl'

# krew (kubectl plugin manager) bin
export PATH="${KREW_ROOT:-$HOME/.krew}/bin:$PATH"

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
# --no-use skips nvm_auto, which was 64% of shell startup (~1.1s) per zprof.
# Sourcing drops from 0.81s to 0.01s.
export NVM_DIR="$HOME/.nvm"
[ -s "/opt/homebrew/opt/nvm/nvm.sh" ] && \. "/opt/homebrew/opt/nvm/nvm.sh" --no-use
[ -s "/opt/homebrew/opt/nvm/etc/bash_completion.d/nvm" ] && \. "/opt/homebrew/opt/nvm/etc/bash_completion.d/nvm"

# --no-use leaves no node on PATH, so put the default version there directly.
# Globs only, no nvm calls and no subprocesses. The `n` qualifier sorts
# numerically so v22.9.0 sorts below v22.14.0.
() {
  local want
  [[ -r $NVM_DIR/alias/default ]] || return
  read -r want < $NVM_DIR/alias/default
  local -a vers=($NVM_DIR/versions/node/v${want#v}*(/Nn))
  (( $#vers )) && export PATH="${vers[-1]}/bin:$PATH"
}

# place this after nvm initialization!
# NB: link/oh-my-zsh/plugins/nvm/nvm.plugin.zsh registers its own load-nvmrc
# chpwd hook. It is not in plugins=(), so it does not load; adding it there
# would give you two competing hooks.
autoload -U add-zsh-hook
# Walks up for .nvmrc with zsh builtins, so a cd into a plain directory costs
# nothing. Only directories that actually pin a version pay for nvm. The old
# version called `nvm version` three times on every single cd (~510ms).
load-nvmrc() {
  local dir=$PWD nvmrc=""
  while [[ -n $dir ]]; do
    if [[ -f $dir/.nvmrc ]]; then
      nvmrc=$dir/.nvmrc
      break
    fi
    dir=${dir%/*}
  done

  if [[ -n $nvmrc ]]; then
    local want
    read -r want < $nvmrc
    if [[ $want != "$_NVM_AUTO_WANT" ]]; then
      # Try first, install on failure. Probing with `nvm version` beforehand
      # cost an extra 338ms for the same answer.
      nvm use "$want" || nvm install
      _NVM_AUTO_WANT=$want
    fi
  elif [[ -n $_NVM_AUTO_WANT ]]; then
    # Only revert if we were the ones who switched away from the default.
    echo "Reverting to nvm default version"
    nvm use default
    unset _NVM_AUTO_WANT
  fi
}
add-zsh-hook chpwd load-nvmrc
load-nvmrc

# The next line updates PATH for the Google Cloud SDK.
if [ -f '/Users/nat/google-cloud-sdk/path.zsh.inc' ]; then source '/Users/nat/google-cloud-sdk/path.zsh.inc'; fi

# The next line enables shell command completion for gcloud.
if [ -f '/Users/nat/google-cloud-sdk/completion.zsh.inc' ]; then source '/Users/nat/google-cloud-sdk/completion.zsh.inc'; fi

# Infra directory switcher
infra() {
  source ~/Projects/dotFiles/bin/infra-script
}

# postgres
export PATH="/opt/homebrew/opt/postgresql@18/bin:$PATH"

# zmv
autoload -Uz zmv

# Wiz
autoload -U +X bashcompinit && bashcompinit
complete -o nospace -C /opt/homebrew/bin/wizcli wizcli

# vim: set filetype=zsh:
