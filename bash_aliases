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

alias vi='vim'

## Extra Cool Alias
alias c="clear"
gfind () { if [ $# -lt 2 ]; then files="*"; search="${1}"; else files="${1}"; search="${2}"; fi; find . -name "$files" -a ! -wholename '*/.*' -exec grep -Hin ${3} "$search" {} \; ; }
pidof () { ps -Acw | egrep -i $@ | awk '{print $1}'; }

## BitTorrent Aliases, require bitornado   
alias btdl="screen btdownloadcurses.bittornado"
alias bthere="screen btlaunchmanycurses.bittornado ."
alias bt="screen rtorrent"

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

## nice for gnome based systems.
alias go='gnome-open'

# Csh compatability:
alias unsetenv=unset
function setenv () {
  export $1="$2"
}

# Socat is awesome.
#if [ -f `which socat 2>/dev/null` ]; then
 #  alias sml="socat READLINE EXEC:sml";
#fi;

# Function which adds an alias to the current shell and to
# the ~/.bash_aliases file.
add-alias () {
   local name=$1 value="$2"
   echo alias $name=\'$value\' >>~/.bash_aliases
   eval alias $name=\'$value\'
   alias $name
}

# "repeat" command.  Like:
#
#       repeat 10 echo foo
repeat () { 
    local count="$1" i;
    shift;
    for i in $(seq 1 "$count");
    do
        eval "$@";
    done
}

# Subfunction needed by Repeat'.
seq () { 
    local lower upper output;
    lower=$1 upper=$2;

    if [ $lower -ge $upper ]; then return; fi
    while [ $lower -le $upper ];
    do
        echo -n "$lower "
        lower=$(($lower + 1))
    done
    echo "$lower"
}

## Aliases added by the system call

