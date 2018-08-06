# ZSH config for work mbp

export TZ=UTC

# Try and fix PATH
export PATH="/usr/local/sbin:/usr/local/bin:$PATH"

# for https://github.com/firstlookmedia/aws-profile-gpg
export AWS_PROFILE_GPG_HOME=$HOME/Projects/aws-profile-gpg
[ -f /usr/local/etc/profile.d/z.sh ] && source /usr/local/etc/profile.d/z.sh
export PATH="/usr/local/opt/gettext/bin:$PATH"

# Colors
export CLICOLOR=1
export LSCOLORS=ExFxCxDxBxegedabagacad

# Golang
export GOPATH=~/Projects/
export PATH="~/Projects/bin:$PATH"

# RVM
[[ -s "$HOME/.rvm/scripts/rvm" ]] && source "$HOME/.rvm/scripts/rvm" # Load RVM into a shell session *as a function*
export PATH="$PATH:$HOME/.rvm/bin"
