/*
Copyright © 2024 NullNUMMER24
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var ProjectFiles string = "." // This defines where all the files are getting stored
var ErrorColor string = "\033[38;5;9m"
var NoColor string = "\033[0m"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "handy-cli",
	Short: "A more or less useless cli",
	Long:  `A more or less useless cli which aims to automate some irrelevant tasks`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
