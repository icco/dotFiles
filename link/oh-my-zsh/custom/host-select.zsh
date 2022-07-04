custom=${ZSH}/custom/hosts

# strip router bullshit
host=$(hostname | sed 's/\./ /g' | awk '{ print $1 }')
if [[ -f $custom/$host.zsh ]]; then
   . $custom/$host.zsh
   echo "===> Loaded $host.zsh";
fi
