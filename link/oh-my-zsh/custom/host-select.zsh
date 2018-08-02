if [ -f ./hosts/`hostname`.zsh ]; then
   . ./hosts/`hostname`.zsh
   echo "===> Loaded `hostname`.zsh";
fi
