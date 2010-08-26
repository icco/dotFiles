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

alias blame-dave="update-dev";

fortune ~/UnWork/crackquotes/crackquotes

# Wrapper around SVN merge and commit for LiveCode
# Stolen from Dave...
lc () {
   revisions=""

   for to in $*
   do
      from=$((to - 1))
      echo "svn merge -r $from:$to file:///var/ifixit/CodeRepos/trunk /home/$USER/LiveCode"
      svn merge -r $from:$to file:///var/ifixit/CodeRepos/trunk /home/$USER/LiveCode
      echo

      revisions="$revisions r$to"
   done

   read -p "Commit? (y/N): " -n 1
   echo
   if [[ $REPLY =~ ^[Yy]$ ]]; then
      echo "svn ci /home/$USER/LiveCode -m \"Merge$revisions to LiveCode\""
      svn ci /home/$USER/LiveCode -m "Merge$revisions to LiveCode"
   fi
}
