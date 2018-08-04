# Nat's zsh Globals

# For building debian packages
export DEBFULLNAME="Nat Welch"
export DEBEMAIL="nat@natwelch.com"
export EDITOR="/usr/bin/vim"
export GREP_COLOR="1;33"
alias grep='grep --color=auto'

# If I have a bin in my user directory, check there for commands.
# We do this late in the file so it takes priority.
[ -d ~/bin ] && export PATH=~/bin:$PATH

if [ -f ~/.mybashrc ]; then
   . ~/.mybashrc
   echo "===> Loaded .mybashrc";
fi

# And we're done!
echo "===> Loaded globals.zsh";
