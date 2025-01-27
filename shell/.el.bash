# store the old dir in a local variable
# reasons:
# `cd folder` must have the old_dir as folder in which the command was launched
old_dir="$(pwd)"

# launch el save to change the el history file
precmd(){
  el save -f "$old_dir" "$(fc -ln -1 | sed 's/^[ \t]*//')"
  old_dir="$(pwd)"
}

# suggestion function, uses fzf to select the command to suggest
function sugg() {
  output=$(el | fzf -q "$READLINE_LINE" +m --read0 | sed 's/\\\\n/\\\n/g')

  READLINE_LINE=${output#*$'\t'}
  if [[ -z "$READLINE_POINT" ]]; then
    echo "$READLINE_LINE"
  else
    READLINE_POINT=0x7fffffff
  fi
}

# bind sugg with Ctrl+Space
bind -m emacs-standard -x '"\C-@": sugg'
bind -m vi-command -x '"\C-@": sugg'
bind -m vi-insert -x '"\C-@": sugg'
