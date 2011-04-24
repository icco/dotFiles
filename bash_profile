# ~/.bash_profile: executed by bash(1) for login shells.
# see /usr/share/doc/bash/examples/startup-files for examples.
# the files are located in the bash-doc package.

TZ='America/Los_Angeles'; export TZ

# the default umask is set in /etc/login.defs
#umask 022

# include .bashrc if it exists
if [ -f ~/.bashrc ]; then 
   . ~/.bashrc 
fi

# set PATH so it includes user's private bin if it exists
if [ -d ~/bin ]; then 
   PATH=~/bin:"${PATH}"; 
fi

# vim: set filetype=sh:
[[ -s "$HOME/.rvm/scripts/rvm" ]] && source "$HOME/.rvm/scripts/rvm" # Load RVM into a shell session *as a function*
