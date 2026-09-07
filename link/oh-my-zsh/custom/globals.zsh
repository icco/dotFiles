# Nat's zsh Globals

# For building debian packages
export DEBFULLNAME="Nat Welch"
export DEBEMAIL="nat@natwelch.com"
export EDITOR="/usr/bin/vim"
export GREP_COLORS="1;33"
alias grep='grep --color=auto'

# If I have a bin in my user directory, check there for commands.
# We do this late in the file so it takes priority.
[ -d ~/bin ] && export PATH=~/bin:$PATH

if [ -f ~/.mybashrc ]; then
   . ~/.mybashrc
   echo "===> Loaded .mybashrc";
fi

# nvm, shared by every host that has it installed.
#
# Sourcing nvm.sh normally runs nvm_auto, which was 64% of a 2.3s zsh startup
# on this laptop (zprof: 1108ms). `--no-use` skips it and drops the source from
# 0.81s to 0.01s, but leaves no node on PATH -- so we resolve the default
# version ourselves with globs and builtins only.
nvm_init_lazy() {
  export NVM_DIR="${NVM_DIR:-$HOME/.nvm}"

  local f
  for f in /opt/homebrew/opt/nvm/nvm.sh "$NVM_DIR/nvm.sh"; do
    [[ -s $f ]] || continue
    \. "$f" --no-use
    break
  done
  (( $+functions[nvm] )) || return 0

  for f in /opt/homebrew/opt/nvm/etc/bash_completion.d/nvm "$NVM_DIR/bash_completion"; do
    [[ -s $f ]] && { \. "$f"; break; }
  done

  nvm_default_bin_to_path
}

# Resolve $NVM_DIR/alias/default to an installed version dir and prepend its
# bin. No nvm calls, no subprocesses.
nvm_default_bin_to_path() {
  local want="default" i
  # Follow up to 5 alias hops (default -> lts/* -> lts/krypton -> v24.20.0).
  # "lts/*" is a literal filename, so every $want expansion must be quoted.
  for i in 1 2 3 4 5; do
    [[ -r "$NVM_DIR/alias/$want" ]] || break
    read -r want < "$NVM_DIR/alias/$want" || break
  done

  local -a vers
  # The `n` qualifier sorts numerically, so v22.9.0 sorts below v22.14.0.
  [[ "${want#v}" == [0-9]* ]] && vers=($NVM_DIR/versions/node/v"${want#v}"*(/Nn))
  # Unresolvable alias (node/stable/system), missing or dead-end alias file, or
  # no matching install: take the newest installed rather than calling nvm,
  # which would either cost the 1.1s back or fail and leave PATH with no node.
  (( $#vers )) || vers=($NVM_DIR/versions/node/v*(/Nn))

  (( $#vers )) && export PATH="${vers[-1]}/bin:$PATH"
}

# And we're done!
echo "===> Loaded globals.zsh";
