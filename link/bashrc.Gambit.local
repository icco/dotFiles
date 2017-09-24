# Bashrc for mac mini

# Try and fix PATH
export PATH="/usr/local/sbin:/usr/local/bin:$PATH"

# Mac ls does not have color option...
alias ls="`which ls`";

# bash completion in osx, thanks to homebrew
if [ -f `brew --prefix`/etc/bash_completion ]; then
  . `brew --prefix`/etc/bash_completion
fi

# Colors
export CLICOLOR=1
export LSCOLORS=ExFxCxDxBxegedabagacad

[[ -s "$HOME/.rvm/scripts/rvm" ]] && source "$HOME/.rvm/scripts/rvm"  # This loads RVM into a shell session.

# For jumping
[ -f /usr/local/etc/profile.d/z.sh ] && source /usr/local/etc/profile.d/z.sh

# vim: set filetype=sh:
