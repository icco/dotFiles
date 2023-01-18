setopt prompt_subst

autoload -U colors && colors

function compute_etu() {
  if type "etu" > /dev/null; then
    echo " ($(etu timesince))"
  fi

  return 0
}

# git_prompt_info is a shell function from /Users/nat/.oh-my-zsh/lib/git.zsh

PROMPT=$'\n[ %{$fg[red]%}%D{%a %b %d %H:%M:%S}%{$reset_color%} ]%{$fg[grey]%}$(compute_etu)%{$reset_color%} $(git_prompt_info) \n[ %b%n@%m %{$fg[blue]%}$(shrink_path -f)%{$reset_color%} ]\\$ '

PS2=$' \e[0;34m%}%B>%{\e[0m%}%b '

# vim: set filetype=zsh:
