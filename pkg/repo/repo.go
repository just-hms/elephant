package repo

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/just-hms/elephant/pkg/entity"
)

type repository struct {
	path string
}

func New(path string) *repository {
	return &repository{
		path: path,
	}
}

// Save writes a command to a file, ensuring that newlines in the command are properly escaped.
func (r *repository) Replace(cmds []entity.Cmd) error {
	destFile, err := os.Create(r.path)
	if err != nil {
		return err
	}
	defer destFile.Close()

	for _, c := range cmds {
		_, err := destFile.WriteString(fmt.Sprintf("%s %s\n", c.Folder, c.Value))
		if err != nil {
			return err
		}
	}

	return nil
}

// Save writes a command to a file, ensuring that newlines in the command are properly escaped.
func (r *repository) Save(c entity.Cmd) error {
	file, err := os.OpenFile(r.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%s %s\n", c.Folder, c.Value))
	return err
}

// load reads commands from a file, applying a filter function to each line.
func (r *repository) load(filter func(e string) bool) ([]entity.Cmd, error) {
	file, err := os.Open(r.path)
	if err != nil {
		// create the file if it doesn't exists
		file, err = os.OpenFile(r.path, os.O_APPEND|os.O_CREATE|os.O_RDONLY, 0644)
		if err != nil {
			return nil, err
		}
	}
	defer file.Close()

	var cmds []entity.Cmd
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		vals := strings.SplitN(line, " ", 2)
		if len(vals) != 2 {
			return nil, errors.New("format error")
		}

		c := entity.Cmd{
			Folder: vals[0],
			Value:  vals[1],
		}

		if filter(c.Folder) {
			cmds = append(cmds, c)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return cmds, nil
}

// LoadAll returns all commands
func (r *repository) LoadAll() ([]entity.Cmd, error) {
	return r.load(func(e string) bool { return true })
}

// LoadAll returns all command inside a folder
func (r *repository) LoadFolder(folder string) ([]entity.Cmd, error) {
	return r.load(func(e string) bool { return e == folder })
}
