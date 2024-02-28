# Store the old dir in a local variable
# reasons:
# `cd folder` must have the old_dir as folder in which the command was launched
old_dir="$(pwd)"

# Function to save the command history using 'el'
save_command_history() {
  el save -f "$old_dir" "$(fc -ln -1)"
  old_dir="$(pwd)"
}

PROMPT_COMMAND="save_command_history"
