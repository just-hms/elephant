package display

import (
	"fmt"

	"github.com/just-hms/elephant/pkg/entity"
)

func PrintValues(cmds []entity.Cmd) {
	for _, c := range cmds {
		fmt.Println(c.Value + "\x00")
	}
}

func PrintFolders(cmds []entity.Cmd) {
	for _, c := range cmds {
		fmt.Println(c.Folder, c.Value+"\x00")
	}
}
