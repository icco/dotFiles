get_pwd() {
  # How many characters of the $PWD should be kept
  local pwdmaxlen=23
  # Indicate that there has been dir truncation
  local trunc_symbol="..."
  local dir=${PWD##*/}
  pwdmaxlen=$(( ( pwdmaxlen < ${#dir} ) ? ${#dir} : pwdmaxlen ))
  local NEW_PWD=${PWD/#$HOME/\~}
  local pwdoffset=$(( ${#NEW_PWD} - pwdmaxlen ))
  if [ ${pwdoffset} -gt "0" ]; then
    NEW_PWD=${NEW_PWD:$pwdoffset:$pwdmaxlen}
    NEW_PWD=${trunc_symbol}/${NEW_PWD#*/}
  fi

  echo $NEW_PWD
}

PROMPT=$'\n[ $fg_bold[red]%D{%a %b %d %H:%M:%S}$reset_color ] $(git_prompt_info) \n[ %b%n@%m $fg_bold[blue]$(get_pwd)$reset_color ]\\$ '

PS2=$' \e[0;34m%}%B>%{\e[0m%}%b '

# vim: set filetype=zsh:
