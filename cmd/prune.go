/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// pruneCmd represents the prune command
var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "remove all the commands that were launched in folder that don't exist anymore (unimplemented)",
	Long:  ``,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		panic("unimplemented")
	},
}

func init() {
	rootCmd.AddCommand(pruneCmd)
}
