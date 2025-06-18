setopt prompt_subst

autoload -U colors && colors
autoload -Uz vcs_info

function compute_etu() {
  if type "etu" > /dev/null; then
    echo " ($(etu timesince))"
  fi

  return 0
}
_omz_register_handler compute_etu

function etu_prompt_status {
  echo -n $_OMZ_ASYNC_OUTPUT[compute_etu]
}

function virtualenv_prompt_info {
  [[ -n ${VIRTUAL_ENV} ]] || return
  echo "${VIRTUAL_ENV:t:gs/%/%%}"
}

zstyle ':vcs_info:*' actionformats \
    '%F{5}(%f%s%F{5})%F{3}-%F{5}[%F{2}%b%F{3}|%F{1}%a%F{5}]%f '
zstyle ':vcs_info:*' formats       \
    '%F{5}(%f%s%F{5})%F{3}-%F{5}[%F{2}%b%F{5}]%f '
zstyle ':vcs_info:(sv[nk]|bzr):*' branchformat '%b%F{1}:%F{3}%r'
precmd () { vcs_info }

# git_prompt_info is a shell function from /Users/nat/.oh-my-zsh/lib/git.zsh
# git_super_status is a shell function from /Users/nat/.oh-my-zsh/plugins/git-prompt/git-prompt.plugin.zsh
# ZSH_THEME_GIT_PROMPT_CACHE=1
ZSH_THEME_GIT_PROMPT_BRANCH=""
ZSH_THEME_GIT_PROMPT_STAGED="%{●%G%}"
ZSH_THEME_GIT_PROMPT_CONFLICTS="%{✖%G%}"
ZSH_THEME_GIT_PROMPT_CHANGED="%{✚%G%}"
ZSH_THEME_GIT_PROMPT_DELETED="%{-%G%}"
ZSH_THEME_GIT_PROMPT_BEHIND="%{↓%G%}"
ZSH_THEME_GIT_PROMPT_AHEAD="%{↑%G%}"
ZSH_THEME_GIT_PROMPT_UNTRACKED="%{…%G%}"
ZSH_THEME_GIT_PROMPT_STASHED="%{⚑%G%}"
ZSH_THEME_GIT_PROMPT_CLEAN="%{✔%G%}"

PROMPT=$'\n[ %{$fg[red]%}%D{%a %b %d %H:%M:%S}%{$reset_color%} ]%{$fg[grey]%}$(etu_prompt_status)%{$reset_color%} ${vcs_info_msg_0_} %{$fg[green]%}$(virtualenv_prompt_info)%{$reset_color%}% \n[ %b%n@%m %{$fg[blue]%}$(shrink_path -f)%{$reset_color%} ]\\$ '
#PROMPT=$'\n[ %{$fg[red]%}%D{%a %b %d %H:%M:%S}%{$reset_color%} ] $(git_prompt_info) \n[ %b%n@%m %{$fg[blue]%}$(shrink_path -f)%{$reset_color%} ]\\$ '

PS2=$' \e[0;34m%}%B>%{\e[0m%}%b '
RPROMPT=''

# vim: set filetype=zsh:
