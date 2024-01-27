/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/just-hms/elephant/internal"
	"github.com/spf13/cobra"
)

var version bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "el",
	Short: "A small program with a long memory",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			fmt.Println(internal.Version)
			return
		}
		showCmd.Run(cmd, args)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "show the version of the program")
	rootCmd.Flags().BoolVarP(&allFlag, "all", "a", false, "show all the history")
	rootCmd.Flags().StringVarP(&showFolderFlag, "folder", "f", "", "show the history of a specified folder")
	rootCmd.MarkFlagsMutuallyExclusive("all", "folder", "version")
}
