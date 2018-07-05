## Handy file for aliases, keep them seperate you here?

# Write vim log to disk
# http://www.drbunsen.org/vim-croquet/
alias vim='vim -w ~/.vimlog "$@"'

## enable color support of ls and also add handy aliases
if [ "$TERM" != "dumb" ]; then
    #eval "dircolors -b"
    alias ls='ls --color=auto'
    alias grep='grep --color=auto'
fi

## some more ls aliases
alias ll='ls -lh'
alias lt='ls -lth'
alias la='ls -A'
alias l='ls -CF'

pidof () { ps -Acw | egrep -i $@ | awk '{print $1}'; }

# A better which
#alias which='alias | /usr/bin/which --tty-only --read-alias --show-dot --show-tilde'

# Job Stuff
alias h='history'
alias j="jobs -l"

# Directory jumping
alias pu="pushd"
alias po="popd"

## duh
alias motd="cat /etc/motd"

# For epic typing fail
alias bim="vim"
alias cim="vim"
alias vi='vim'

# Csh compatability:
alias unsetenv=unset
function setenv () {
  export $1="$2"
}

# free env
dotenv () {
  set -a
  [ -f .env ] && . .env
  set +a
}

alias fuck='sudo $(history -p \!\!)'

# Generates an aws_session so you can role switch.
function aws_session() {
  # Clear out existing AWS session environment, or the awscli call will fail
  unset AWS_ACCESS_KEY_ID AWS_SECRET_ACCESS_KEY AWS_SESSION_TOKEN AWS_SECURITY_TOKEN
  # Old ec2 tools use other env vars
  unset AWS_ACCESS_KEY AWS_SECRET_KEY AWS_DELEGATION_TOKEN

  if [ "$#" -eq 0 ]; then
    echo "Usage: aws_session profile_name"
    echo ""
    echo "Creates a session for the specified profile. The profile must be"
    echo "defined in ~/.aws/config."
    echo
    echo "Suggested config:"
    echo
    echo "[profile default]"
    echo "output = json"
    echo "region = us-east-1"
    echo
    echo "[profile ops-admin]"
    echo "role_arn = arn:aws:iam::123456789123:role/role-name"
    echo "source_profile = default"
    echo
    echo "And in ~/.aws/credentials"
    echo
    echo "[default]"
    echo "aws_access_key_id = AKIABLAHBLAHBLAH1234"
    echo "aws_secret_access_key = asdjhlkjhljkhlsajkdhjkldhajhlkajhdkjlshd"
    echo
  else
    local profile=$1
    local username=$(aws iam get-user --profile=$profile --query User.UserName --output text)
    local device=$(aws iam list-mfa-devices --profile=$profile --user-name $username --query MFADevices[0].SerialNumber --output text)
    echo -n "Authentication code for ${device} (${profile}/${username}): "

    local code
    read code

    local sts=$(aws sts get-session-token --profile=$profile --duration-seconds 3600 --serial-number $device --token-code $code --output text)

    export AWS_ACCESS_KEY_ID=$(echo "$sts" | cut -f 2)
    export AWS_SECRET_ACCESS_KEY=$(echo "$sts" | cut -f 4)
    export AWS_SESSION_TOKEN=$(echo "$sts" | cut -f 5)
  fi
}
alias aws-session=aws_session

function aws-assume-role() {
  if [ "$#" -eq 0 ]; then
    echo "TODO: Write usage"
  else
    local profile=$1
    local role=$2
    aws-session $profile

    local sts=$(aws sts assume-role --role-arn $role --role-session-name $USER)

    export AWS_ACCESS_KEY_ID=$(echo "$sts" | jq -r .Credentials.AccessKeyId)
    export AWS_SECRET_ACCESS_KEY=$(echo "$sts" | jq -r .Credentials.SecretAccessKey)
    export AWS_SESSION_TOKEN=$(echo "$sts" | jq -r .Credentials.SessionToken)

    echo $sts | jq .AssumedRoleUser
  fi
}

# vim: set filetype=sh:
