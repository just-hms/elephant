/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/just-hms/elephant/internal"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "display the version of el",
	Long:  `display the version of els`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(internal.Version)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
