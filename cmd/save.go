/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"path"
	"path/filepath"

	"github.com/just-hms/elephant/pkg/entity"
	"github.com/just-hms/elephant/pkg/repo"
	"github.com/spf13/cobra"
)

var saveFolderFlag string

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "save a command specifying a folder and a value",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		h, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		r := repo.New(path.Join(h, ".history.el"))

		pat, err := filepath.Abs(saveFolderFlag)
		if err != nil {
			return err
		}
		err = r.Save(entity.Cmd{
			Folder: pat,
			Value:  args[0],
		})
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
	saveCmd.Flags().StringVarP(&saveFolderFlag, "folder", "f", "", "show the history of a specified folder")
}
