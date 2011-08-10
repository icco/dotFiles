# Bashrc for my work laptop
export CC=gcc-4.2

# Mac ls does not have color...
alias ls="`which ls`";

if [ -f /opt/local/etc/bash_completion ]; then
    . /opt/local/etc/bash_completion
fi

# Colors
export CLICOLOR=1
export LSCOLORS=ExFxCxDxBxegedabagacad

#`brew --prefix grc`/etc/grc.bashrc

# vim: set filetype=sh:
