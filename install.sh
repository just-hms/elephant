#!/bin/bash

zsh_file="$HOME/.zshrc"
zsh_sourcer="[ -f ~/.el.zsh ] && source ~/.el.zsh"

touch $HOME/.history.el

if [ -f $zsh_file ] ; then
    
     
    
    if ! grep -q "$zsh_sourcer" $HOME/.zshrc ; then
        echo "" >> $zsh_file
        echo "# el" >> $zsh_file
        echo $zsh_sourcer >> $zsh_file
    else
        echo "el.zsh was already sourced in \"$zsh_file\""
        echo "updating..."
    fi;
    
    cp shell/el.zsh $HOME/.el.zsh

    echo "> source \"$zsh_file\""
fi
