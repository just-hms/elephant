/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/just-hms/elephant/pkg/shell"
	"github.com/spf13/cobra"
)

var (
	zshFlag  bool
	bashFlag bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "el",
	Short: "A small program with a long memory",
	RunE: func(cmd *cobra.Command, args []string) error {
		if zshFlag {
			fmt.Println(shell.ZSH)
			return nil
		}
		if bashFlag {
			fmt.Println(shell.ZSH)
			return nil
		}
		return showCmd.RunE(cmd, args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVar(&zshFlag, "zsh", false, "pass this to output the zsh key-bindings file")
	rootCmd.Flags().BoolVar(&bashFlag, "bash", false, "pass this to output the bash key-bindings file")
}
