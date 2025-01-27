<p align="center">
    <img style="width:8em;" src="./assets/logo.png" alt="elephant-logo">
</p>

# Elephant

> A small program with a long memory

- run `el` to show you the list of command launched in the current folder.
- use <kbd>Ctrl</kbd>+<kbd>Spacebar</kbd> to show you suggestions based on the history in the current folder. _(Requires `fzf`)_

## Install

1. *Install `el`*:

```sh
go install github.com/just-hms/elephant@latest
```

Alternatively, download the precompiled binary from the [latest](https://github.com/just-hms/elephant/releases/latest) release and move it to your system's `PATH`

---

2. *Set up shell integration*

Add the following line to your shell configuration file:

- bash

  ```sh
  eval "$(el --bash)"
  ```

- zsh

  ```sh
  source <(el --zsh)
  ```

---

3. (OPTIONAL) Install https://github.com/junegunn/fzf to use keybindings

## Getting Started

After installation, Elephant is immediately ready to use. Begin by navigating to any directory and executing commands as you normally would. Invoke Elephant by typing `el` or use the <kbd>Ctrl</kbd>+<kbd>Spacebar</kbd> shortcut for suggestions on what to run next.

# History file

Elephant stores its history in a plain-text file located at: `$HOME/.history.el`.

This file is human-readable and editable, giving you full control over your command history.
