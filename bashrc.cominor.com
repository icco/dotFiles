# .bashrc

# Source global definitions
if [ -f /etc/bashrc ]; then
	. /etc/bashrc
fi

# User specific aliases and functions
PATH=/usr/local/bin:/usr/local/mysql/bin/:/usr/local/sbin/:/usr/sbin/:/sbin/:$PATH:$HOME/bin

export EDITOR=vim

# Set up PHP and code paths for iFixit CLI scripts and phing
export MY_CODE_DIR=/home/nwelch/Code
export PHP_CLASSPATH=${PHING_HOME}/classes:${MY_CODE_DIR}:${MY_CODE_DIR}/3P

function mgrep() { grep -r \'$1\' . | grep -v svn ; }
