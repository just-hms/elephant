package shell

import _ "embed"

//go:embed .el.zsh
var ZSH string

//go:embed .el.bash
var BASH string
