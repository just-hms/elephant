package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/just-hms/elephant/pkg/display"
	"github.com/just-hms/elephant/pkg/entity"
	"github.com/just-hms/elephant/pkg/repo"
)

var flagSave = flag.Bool("save", false, "save a command specifiying a folder and a value")
var flagFolder = flag.String("folder", "", "specify a folder")
var flagAll = flag.Bool("all", false, "similar to `history`")

func CurrentFolder() string {
	path, _ := os.Getwd()
	return path
}

func main() {
	flag.Parse()
	args := flag.Args()

	h, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	r := repo.Repo{
		Path: path.Join(h, "history.el"),
	}

	if *flagSave {
		if len(args) != 2 {
			panic(fmt.Sprintf("[save] wrong arguments %d", len(args)))

		}
		pat, err := filepath.Abs(args[0])
		if err != nil {
			panic(err)
		}
		err = r.Save(&entity.Cmd{
			Folder: pat,
			Value:  args[1],
		})
		if err != nil {
			panic(err)
		}
		return
	}
	if *flagAll {
		if len(args) != 0 {
			panic(fmt.Sprintf("[all] wrong arguments %d", len(args)))
		}
		cmds, err := r.LoadAll()
		if err != nil {
			panic(err)
		}
		display.PrintHistory(cmds)
		return
	}
	if *flagFolder != "" {
		if len(args) != 0 {
			panic(fmt.Sprintf("[folder] wrong arguments %d", len(args)))
		}
		cmds, err := r.LoadFolder(*flagFolder)
		if err != nil {
			panic(err)
		}
		display.Println(cmds)
		return
	}

	if len(args) != 0 {
		panic(fmt.Sprintf("[def] wrong arguments %d", len(args)))
	}

	pat, err := filepath.Abs(CurrentFolder())
	if err != nil {
		panic(err)
	}

	cmds, err := r.LoadFolder(pat)
	if err != nil {
		panic(err)
	}
	display.Println(cmds)
}
