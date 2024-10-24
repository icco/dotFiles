setopt prompt_subst

autoload -U colors && colors

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

# git_prompt_info is a shell function from /Users/nat/.oh-my-zsh/lib/git.zsh

PROMPT=$'\n[ %{$fg[red]%}%D{%a %b %d %H:%M:%S}%{$reset_color%} ]%{$fg[grey]%}$(etu_prompt_status)%{$reset_color%} $(git_prompt_info) %{$fg[green]%}$(virtualenv_prompt_info)%{$reset_color%}% \n[ %b%n@%m %{$fg[blue]%}$(shrink_path -f)%{$reset_color%} ]\\$ '
#PROMPT=$'\n[ %{$fg[red]%}%D{%a %b %d %H:%M:%S}%{$reset_color%} ] $(git_prompt_info) \n[ %b%n@%m %{$fg[blue]%}$(shrink_path -f)%{$reset_color%} ]\\$ '

PS2=$' \e[0;34m%}%B>%{\e[0m%}%b '
RPROMPT=''

# vim: set filetype=zsh:
