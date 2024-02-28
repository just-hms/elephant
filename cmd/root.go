/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "el",
	Short: "A small program with a long memory",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
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
	rootCmd.Flags().BoolVarP(&allFlag, "all", "a", false, "show all the history")
	rootCmd.Flags().StringVarP(&showFolderFlag, "folder", "f", "", "show the history of a specified folder")
	rootCmd.MarkFlagsMutuallyExclusive("all", "folder")
}
