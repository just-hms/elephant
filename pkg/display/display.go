package display

import (
	"fmt"

	"github.com/just-hms/elephant/pkg/entity"
)

func PrintValues(cmds []entity.Cmd) {
	for _, c := range cmds {
		fmt.Println(c.Value)
	}
}

func PrintFolders(cmds []entity.Cmd) {
	for _, c := range cmds {
		fmt.Println(c.Folder, c.Value)
	}
}
