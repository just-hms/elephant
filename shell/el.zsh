old_dir="$(pwd)"
precmd(){ 
  el --save "$old_dir" "$(fc -ln -1)"
  old_dir="$(pwd)"
}

# write el.. to receive suggestions
function sugg() {
    if [[ $BUFFER == "el." ]]; then
        BUFFER=`el | sort -u | fzf`
    else
      BUFFER="$BUFFER."
    fi
    CURSOR=$#BUFFER  # Move cursor to the end of the buffer
}

zle -N sugg
bindkey '.' sugg
