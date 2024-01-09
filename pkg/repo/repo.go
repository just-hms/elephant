package repo

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/just-hms/elephant/pkg/entity"
)

type Repo struct {
	Path string
}

// Save writes a command to a file, ensuring that newlines in the command are properly escaped.
func (r *Repo) Save(c *entity.Cmd) error {
	file, err := os.OpenFile(r.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	escapedFolder := strings.ReplaceAll(c.Folder, "\n", "\\n")
	escapedValue := strings.ReplaceAll(c.Value, "\n", "\\n")
	_, err = file.WriteString(fmt.Sprintf("%s %s\n", escapedFolder, escapedValue))
	return err
}

// load reads commands from a file, applying a filter function to each line.
func (r *Repo) load(filter func(e string) bool) ([]entity.Cmd, error) {
	file, err := os.Open(r.Path)
	if err != nil {
		return nil, err
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
func (r *Repo) LoadAll() ([]entity.Cmd, error) {
	return r.load(func(e string) bool { return true })
}

// LoadAll returns all command inside a folder
func (r *Repo) LoadFolder(folder string) ([]entity.Cmd, error) {
	return r.load(func(e string) bool { return e == folder })
}
