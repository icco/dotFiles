# tirith - Terminal security
# https://github.com/sheeki03/tirith

if (( ! $+commands[tirith] )); then
  # TIRITH_STRICT=1: exit 1 kills shell startup (use with caution)
  # Default: warn and continue (safe for plugin-first installs)
  if [[ "${TIRITH_STRICT:-0}" == "1" ]]; then
    echo "[tirith] FATAL: tirith not found. Install first, then restart." >&2
    exit 1  # Hard kill - shell won't start
  fi
  echo "[tirith] Warning: tirith not found. Install: brew install sheeki03/tap/tirith" >&2
  return
fi

eval "$(tirith init)"
