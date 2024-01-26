old_dir="$(pwd)"
precmd(){ 
  el save -f "$old_dir" "$(fc -ln -1)"
  old_dir="$(pwd)"
}

# write el.. to receive suggestions
function sugg() {
    # Get the left part of the buffer up to the cursor position
    local left_part=${BUFFER:0:$CURSOR}
    
    if [[ -n $left_part ]]; then        
      BUFFER=$(el | sort -u | fzf -q "$left_part")
    else;
      BUFFER=$(el | sort -u | fzf)
    fi
    
    # Update the cursor position
    CURSOR=$#BUFFER
}

# bind sugg with Shift+Tab
zle -N sugg
bindkey '^[[Z' sugg
