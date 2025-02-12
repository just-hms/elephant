# store the old dir in a local variable
# reasons:
# `cd folder` must have the old_dir as folder in which the command was launched
old_dir="$(pwd)"

# launch el save to change the el history file
precmd(){
  el save -f "$old_dir" "$(fc -ln -1)"
  old_dir="$(pwd)"
}

# suggestion function, uses fzf to select the command to suggest
function sugg() {
  # Get the left part of the buffer up to the cursor position
  local left_part=${BUFFER:0:$CURSOR}

  # change the buffer to match the selection
  BUFFER=$(el | fzf -q "$left_part" | sed 's/\\\\n/\\\n/g')
  # Update the cursor position
  CURSOR=$#BUFFER
}

# bind sugg with Ctrl + Space
zle -N sugg
bindkey '^@' sugg
