# Elephant

> A small program with a log memory

run `el` to show you the list of command launched in the current dir.

works very well with `fzf`


## todo

- [ ] `install.sh` make it installable with a single command
- [ ] add something similar to ctrl+R in bash (even if it is not a shortcut, find a way to write in the curren command)
- [ ] `--purge` removes history of folder which don't exist anymore


## reminder

must paired with inside your `~/.zshrc`

```shell
old_dir="$(pwd)"
precmd(){ 
  el --save "$old_dir" "$(fc -ln -1)"
  old_dir="$(pwd)"
}
```