setopt prompt_subst

ETU=""
if type "etu" > /dev/null; then
  ETU=" ($(etu timesince))"
fi

GIT=""
if type "git_prompt_info" > /dev/null; then
  GIT=" $(git_prompt_info)"
fi

PROMPT=$'\n[ %{$fg[red]%}%D{%a %b %d %H:%M:%S}%{$reset_color%} ]%{$fg[grey]%}$ETU%{$reset_color%}$GIT\n[ %b%n@%m %{$fg[blue]%}$(shrink_path -f)%{$reset_color%} ]\\$ '

PS2=$' \e[0;34m%}%B>%{\e[0m%}%b '

# vim: set filetype=zsh:
