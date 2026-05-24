custom=${ZSH}/custom/hosts

# strip router bullshit
host=$(hostname | sed 's/\./ /g' | awk '{ print $1 }')

# Termux on Android reports "localhost"; ask the system for the real name.
if [[ "$host" == "localhost" ]] && (( $+commands[getprop] )); then
   real=$(getprop net.hostname 2>/dev/null)
   [[ -z "$real" ]] && real=$(getprop ro.product.model 2>/dev/null | tr ' ' '_')
   [[ -n "$real" ]] && host=$real
fi

if [[ -f $custom/$host.zsh ]]; then
   . $custom/$host.zsh
   echo "===> Loaded $host.zsh";
fi
