<p align="center">
    <img style="width:8em;" src="./assets/logo.png" alt="jim">
</p>

# Elephant

> A small program with a long memory

- run `el` to show you the list of command launched in the current folder.
- use <kbd>Ctrl</kbd>+<kbd>Spacebar</kbd> to show you suggestions based on the history in the current folder (requires  `fzf`)

## Install

Install Elephant with a single command:

```sh
go install github.com/just-hms/elephant/el@latest
```

Alternatively, download the precompiled binary from the [latest](https://github.com/just-hms/elephant/releases/latest) release and move it to your system's `PATH`

## Setting up shell integration

> Install https://github.com/junegunn/fzf to use keybindings

Add the following line to your shell configuration file.

- bash

  ```sh
  eval "$(el --bash)"
  ```

- zsh

  ```sh
  source <(el --zsh)
  ```

## Getting Started

After installation, Elephant is immediately ready to use. Begin by navigating to any directory and executing commands as you normally would. Invoke Elephant by typing `el` or use the <kbd>Ctrl</kbd>+<kbd>Spacebar</kbd> shortcut for suggestions on what to run next.
