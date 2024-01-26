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

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "save a command specifiying a folder and a value",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		h, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		r := repo.New(path.Join(h, ".history.el"))

		pat, err := filepath.Abs(folderFlag)
		if err != nil {
			panic(err)
		}
		err = r.Save(&entity.Cmd{
			Folder: pat,
			Value:  args[0],
		})
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
	saveCmd.Flags().StringVarP(&folderFlag, "folder", "f", "", "Show the history of a specified folder")
}
