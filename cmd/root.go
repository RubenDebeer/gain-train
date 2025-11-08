/*
Copyright © 2025 Ruben De Beer <Rubendebeercoding@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gain-train",
	Short: "Your personal fitness companion for tracking workouts and achieving training goals",
	Long: `Gain Train is a powerful command-line fitness tracker that helps you:

	Create and manage custom workout routines
	Track your exercise progress over time  
	Build consistent training habits
	Stay accountable to your fitness goals

	Whether you're a beginner starting your fitness journey or an experienced athlete 
	optimizing your training, Gain Train provides the tools you need to log workouts,
	monitor your progress, and achieve your personal fitness milestones. Awe get on 
	the Gain Train Choo Choo......

	Get started by creating your first workout:
	gain-train workouts create

	View all available commands:
	gain-train help`,
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
