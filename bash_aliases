## Handy file for aliases, keep them seperate you here?
## Word to your mother

## enable color support of ls and also add handy aliases
if [ "$TERM" != "dumb" ]; then
    #eval "dircolors -b"
    alias ls='ls --color=auto'
    alias grep='grep --color=auto'
    alias dir='ls --color=auto --format=vertical'
    alias vdir='ls --color=auto --format=long'   
fi
  
## some more ls aliases
alias ll='ls -lh'
alias la='ls -A'
alias l='ls -CF'

## Extra Cool Alias
alias c="clear"

gfind () { 
   if [ $# -lt 2 ]; then 
      files="*"; 
      search="${1}"; 
   else 
      files="${1}"; 
      search="${2}"; 
   fi; 
   find . -name "$files" -a ! -wholename '*/.*' -exec grep -Hin ${3} "$search" {} \; ; 
}

pidof () { ps -Acw | egrep -i $@ | awk '{print $1}'; }

## IP Address
#alias echoIP="/sbin/ifconfig eth0 | grep 'inet addr:' | cut -d: -f2 | cut -d' ' -f1"
alias myip="curl -s http://natwelch.com/ip/ | html2text"

########### Debian Based Aliases

alias texclean='rm -vf *.toc *.aux *.log *.cp *.fn *.tp *.vr *.pg *.ky'
alias clean='echo -n "Really clean this directory?";
        read yorn;
        if test "$yorn" = "y"; then
           rm -f \#* *~ .*~ *.bak .*.bak  *.tmp .*.tmp core a.out;
           echo "Cleaned.";
        else
           echo "Not cleaned.";
        fi';

# Job Stuff
alias h='history'
alias j="jobs -l"
alias pu="pushd"
alias po="popd"

alias nat="echo 'I\'m the proud owner of DF 82 8A F5 6F EF 0F 15 F6 12 09 0B 03 21 DA CF'"

## duh
alias motd="cat /etc/motd"

## For git repos
alias add_gitignore="echo '*swp' > .gitignore && git add .gitignore && git commit -m 'adds gitignore'"

# For epic typing fail
alias bim="vim"
alias cim="vim"
alias vi='vim'

## nice for gnome based systems.
alias go='gnome-open'

# Csh compatability:
alias unsetenv=unset
function setenv () {
  export $1="$2"
}

# vim: set filetype=sh:
