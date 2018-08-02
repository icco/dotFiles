custom=${ZSH}/custom/hosts
if [[ -f $custom/`hostname`.zsh ]]; then
   . $custom/`hostname`.zsh
   echo "===> Loaded `hostname`.zsh";
fi
