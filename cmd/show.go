/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"path"
	"path/filepath"

	"github.com/just-hms/elephant/pkg/display"
	"github.com/just-hms/elephant/pkg/osx"
	"github.com/just-hms/elephant/pkg/repo"
	"github.com/spf13/cobra"
)

var (
	allFlag        bool
	showFolderFlag string
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show the history",
	Long: `Show provides a detailed view of the command history.

It can display the history of:
- the current folder
- all the history
- the history for a specific folder.
`,
	Args: cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		h, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		r := repo.New(path.Join(h, ".history.el"))

		pat, err := filepath.Abs(osx.CurrentFolder())
		if err != nil {
			return err
		}

		if allFlag {
			cmds, err := r.LoadAll()
			if err != nil {
				panic(err)
			}
			display.PrintFolders(cmds)
			return nil
		}

		if showFolderFlag != "" {
			cmds, err := r.LoadFolder(showFolderFlag)
			if err != nil {
				panic(err)
			}
			display.PrintValues(cmds)
			return nil
		}

		// Default behavior
		cmds, err := r.LoadFolder(pat)
		if err != nil {
			return err
		}
		display.PrintValues(cmds)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	showCmd.Flags().BoolVarP(&allFlag, "all", "a", false, "show all the history")
	showCmd.Flags().StringVarP(&showFolderFlag, "folder", "f", "", "show the history of a specified folder")
	showCmd.MarkFlagsMutuallyExclusive("all", "folder")
}
