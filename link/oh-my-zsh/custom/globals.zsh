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

# git-prompt: prefer the icco/gitstatus binary over the plugin's python script.
#
# Redefined here rather than edited in link/oh-my-zsh/plugins/git-prompt/,
# because that path is vendored oh-my-zsh and `task omz` overwrites it.
# custom/*.zsh is sourced after plugins (oh-my-zsh.sh:221 vs :216), so this
# definition wins. The python fallback keeps the prompt working on every host
# that has not run `brew install icco/tap/gitstatus`.
#
# If a future oh-my-zsh adds fields to this function, this copy goes stale and
# the prompt loses the new field -- it does not break.
#
# The plugin refreshes three times per cd and twice per prompt, each one
# forking git: a chpwd hook, a precmd hook, and a third call from inside
# git_super_status during prompt expansion. Two of those are redundant --
# precmd always runs before the prompt is expanded, so the chpwd hook and the
# git_super_status call have nothing new to find.
#
# Dropping the chpwd hook and memoising for the rest of the cycle leaves
# exactly one refresh per prompt, with no loss of freshness: the memo is
# invalidated at the front of precmd, so every prompt -- including a bare
# Enter -- refreshes once. ZSH_THEME_GIT_PROMPT_CACHE would be cheaper still
# but stops refreshing after non-git commands, leaving the prompt wrong until
# you cd or run git.
if (( $+functions[update_current_git_vars] )); then
  autoload -U add-zsh-hook

  # Redundant with precmd, which runs before every prompt expansion.
  add-zsh-hook -d chpwd chpwd_update_git_vars

  _git_prompt_memo_pwd=""
  _git_prompt_invalidate() { _git_prompt_memo_pwd="" }
  # Must run before precmd_update_git_vars, which the plugin already
  # appended, so prepend rather than add-zsh-hook.
  precmd_functions=(_git_prompt_invalidate ${precmd_functions:#_git_prompt_invalidate})

  update_current_git_vars() {
    # Already refreshed for this directory during this prompt cycle.
    [[ -n $_git_prompt_memo_pwd && $_git_prompt_memo_pwd == $PWD ]] && return
    _git_prompt_memo_pwd=$PWD

    unset __CURRENT_GIT_STATUS

    local _GIT_STATUS
    if (( $+commands[gitstatus] )); then
      _GIT_STATUS=$(gitstatus 2>/dev/null)
    else
      _GIT_STATUS=$(python3 "$__GIT_PROMPT_DIR/gitstatus.py" 2>/dev/null)
    fi
    __CURRENT_GIT_STATUS=("${(@s: :)_GIT_STATUS}")

    GIT_BRANCH=$__CURRENT_GIT_STATUS[1]
    GIT_AHEAD=$__CURRENT_GIT_STATUS[2]
    GIT_BEHIND=$__CURRENT_GIT_STATUS[3]
    GIT_STAGED=$__CURRENT_GIT_STATUS[4]
    GIT_CONFLICTS=$__CURRENT_GIT_STATUS[5]
    GIT_CHANGED=$__CURRENT_GIT_STATUS[6]
    GIT_UNTRACKED=$__CURRENT_GIT_STATUS[7]
    GIT_STASHED=$__CURRENT_GIT_STATUS[8]
    GIT_CLEAN=$__CURRENT_GIT_STATUS[9]
    GIT_DELETED=$__CURRENT_GIT_STATUS[10]

    if [ -z ${ZSH_THEME_GIT_SHOW_UPSTREAM+x} ]; then
      GIT_UPSTREAM=
    else
      GIT_UPSTREAM=$(git rev-parse --abbrev-ref --symbolic-full-name "@{upstream}" 2>/dev/null) && GIT_UPSTREAM="${ZSH_THEME_GIT_PROMPT_UPSTREAM_SEPARATOR}${GIT_UPSTREAM}"
    fi
  }
fi

# And we're done!
echo "===> Loaded globals.zsh";
