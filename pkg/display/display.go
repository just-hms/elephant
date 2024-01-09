package display

import (
	"fmt"

	"github.com/just-hms/elephant/pkg/entity"
)

func Println(cmds []entity.Cmd) {
	for _, c := range cmds {
		fmt.Println(c.Value)
	}
}