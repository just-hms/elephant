/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"path"

	"github.com/just-hms/elephant/pkg/entity"
	"github.com/just-hms/elephant/pkg/repo"
	"github.com/spf13/cobra"
)

// pruneCmd represents the prune command
var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "remove all the commands that were launched in folder that don't exist anymore",
	Long:  ``,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {
		h, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		r := repo.New(path.Join(h, ".history.el"))

		cmds, err := r.LoadAll()
		if err != nil {
			panic(err)
		}

		pruned := []entity.Cmd{}
		cache := map[string]bool{}

		for _, c := range cmds {
			if exists, ok := cache[c.Folder]; ok {
				if exists {
					pruned = append(pruned, c)
				}
				continue
			}
			// if the file exists add it to the cache and add it to the pruned ones
			if _, err := os.Stat(c.Folder); !os.IsNotExist(err) {
				cache[c.Folder] = true
				pruned = append(pruned, c)
			} else {
				cache[c.Folder] = false
			}
		}

		err = r.Replace(pruned)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(pruneCmd)
}
