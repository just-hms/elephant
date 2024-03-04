# store the old dir in a local variable
# reasons:
# `cd folder` must have the old_dir as folder in which the command was launched
old_dir="$(pwd)"

# launch el save to change the el history file
precmd(){
  el save -f "$old_dir" "$(fc -ln -1 | sed 's/^[ \t]*//')"
  old_dir="$(pwd)"
}
