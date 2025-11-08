/*
Copyright © 2025 Ruben De Beer <Rubendebeercoding@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gain-train",
	Short: "A command-line tool to log and track your workouts, monitor progress over time, and keep you accountable to your training goals  ",
	Long:  `I Don't know I will need to thing about it...`,
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
