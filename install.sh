#!/bin/bash

zsh_file="$HOME/.zshrc"
zsh_sourcer="[ -f ~/.el.zsh ] && source ~/.el.zsh"
if [ -f $zsh_file ] ; then
    
    cp shell/el.zsh $HOME/.el.zsh
    echo "el.zsh updated"
    
    if ! grep -q "$zsh_sourcer" $HOME/.zshrc ; then
        echo "" >> $zsh_file
        echo "# el" >> $zsh_file
        echo $zsh_sourcer >> $zsh_file
        echo "source \"$zsh_file\""
    else
        echo "el.zsh was already sourced in \"$zsh_file\""
    fi;
fi
