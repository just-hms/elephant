package display

import (
	"fmt"

	"github.com/just-hms/elephant/pkg/entity"
)

func PrintValues(cmds []entity.Cmd) {
	for _, c := range cmds {
		fmt.Printf("%s\x00", c.Value)
	}
}

func PrintFolders(cmds []entity.Cmd) {
	for _, c := range cmds {
		fmt.Printf("%s %s\x00", c.Folder, c.Value)
	}
}
