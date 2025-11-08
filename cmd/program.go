/*
Copyright © 2025 Ruben De Beer <Rubendebeercoding@gmail.com>
*/
package cmd

import (
	"gain-train-cli/pkg/ui"

	"github.com/spf13/cobra"
)

var programCmd = &cobra.Command{
	Use:     "program",
	Aliases: []string{"p"},
	Short:   "Manage your workout programs",
	Long: `Create, view, modify, and organize your workout programs.

	The program command provides complete management of your workout programs,
	allowing you to build custom routines, track your exercises, and maintain
	a consistent training schedule.`,
}

var createProgramCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	Short:   "Create a New Workout Program",
	Long:    "Create a New Workout Program",
	Run:     createProgramHandler,
}

func init() {
	rootCmd.AddCommand(programCmd)
	initProgramCmd()
}

func initProgramCmd() {
	programCmd.AddCommand(createProgramCmd)
}

func createProgramHandler(cmd *cobra.Command, args []string) {

	// Tell The User we are about to create a Program

	// Begin Prompt
	// Ask a Question and wait for the awnser
	// Name of the Program = User Input
	// Description = User Input
	// workouts = user input or selection

	ui.PrintInfo("Yes Bro let's Create ")
}
