# Nat's Bashrc

# If not running interactively, don't do anything
[ -z "$PS1" ] && return

# If I have a bin in my user directory, check there for commands.
[ -d ~/bin ] && PATH=~/bin:$PATH

# don't put duplicate lines in the history. See bash(1) for more options
export HISTCONTROL=ignoredups

# give me the correct compiler
export CC=/usr/bin/gcc

# check the window size after each command and, if necessary,
# update the values of LINES and COLUMNS.
shopt -s checkwinsize

# make less more friendly for non-text input files, see lesspipe(1)
[ -x /usr/bin/lesspipe ] && eval "$(lesspipe)"

# set variable identifying the chroot you work in (used in the prompt below)
if [ -z "$debian_chroot" ] && [ -r /etc/debian_chroot ]; then
    debian_chroot=$(cat /etc/debian_chroot)
fi

export EDITOR="/usr/bin/vim"

export GREP_COLOR="1;33"
alias grep='grep --color=auto'

# Comment in the above and uncomment this below for a color prompt
#PS1='${debian_chroot:+($debian_chroot)}\[\033[01;32m\]\u@\h\[\033[00m\]:\[\033[01;34m\]\w\[\033[00m\]\$ '

## PS1
# I put PS1 in a sepeerate file
if [ -f ~/.bash_PS1 ]; then
    . ~/.bash_PS1
fi

# Alias definitions.
# You may want to put all your additions into a separate file like
# ~/.bash_aliases, instead of adding them here directly.
# See /usr/share/doc/bash-doc/examples in the bash-doc package.

if [ -f ~/.bash_aliases ]; then
    . ~/.bash_aliases
fi

# enable programmable completion features. Since we don't know if the machine
# supports it, make sure to check both locally and in /etc. System settings
# always take preference.
if [ -f /etc/bash_completion ]; then
   . /etc/bash_completion
elif [ -f ~/.bash_completion ]; then
   . ~/.bash_completion
fi

# For certain machines add an additional bashrc
if [ -f ~/.bashrc.`hostname` ]; then
   . ~/.bashrc.`hostname`
elif [ -f ~/.mybashrc ]; then
   . ~/.mybashrc
fi

# see /usr/share/doc/bash/examples/startup-files for examples

# vim: set filetype=sh:
