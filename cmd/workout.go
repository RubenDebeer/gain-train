/*
Copyright © 2025 Ruben De Beer <Rubendebeercoding@gmail.com>
*/
package cmd

import (
	"fmt"
	"gain-train-cli/pkg/ui"
	"gain-train-cli/pkg/workouts"

	"github.com/spf13/cobra"
)

var workoutsCmd = &cobra.Command{
	Use:     "workouts",
	Aliases: []string{"w"},
	Short:   "Manage your workout routines and training sessions",
	Long: `Create, view, modify, and organize your workout routines.

	The workouts command provides complete management of your training sessions,
	allowing you to build custom routines, track your exercises, and maintain
	a consistent training schedule.

	Examples:
	gain-train workouts add     # Create a new workout routine
	gain-train workouts list    # View all your workouts
	gain-train workouts remove  # Delete a workout routine`,
}

var createWorkoutCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new workout",
	Long:  "Create a new workout to your training plan.",
	Run:   createWorkoutHandler,
}

var listWorkoutsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all workouts",
	Long:  "Display a list of all your workouts.",
	Run:   listWorkoutsHandler,
}

var updateWorkoutCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a workout",
	Long:  "Update an existing workout.",
	Run:   updateWorkoutHandler,
}

var removeWorkoutCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a workout",
	Long:  "Remove a workout from your training plan.",
	Run:   removeWorkoutHandler,
}

func init() {
	rootCmd.AddCommand(workoutsCmd)
	initWorkoutCmd()
}

func initWorkoutCmd() {
	workoutsCmd.AddCommand(createWorkoutCmd)
	workoutsCmd.AddCommand(listWorkoutsCmd)
	workoutsCmd.AddCommand(updateWorkoutCmd)
	workoutsCmd.AddCommand(removeWorkoutCmd)
}

func createWorkoutHandler(cmd *cobra.Command, args []string) {

	ui.PrintInfo("Creating a new workout...")

	storage := workouts.NewJSONStorage()

	name := ui.PromptString("Enter workout name")

	description := ui.PromptString("Enter workout description (optional)")

	workout := &workouts.Workout{
		Name:        name,
		Description: description,
		Exercises:   []workouts.Exercise{},
	}

	for {
		if ui.PromptConfirm("Add an exercise to this workout?") {
			exercise := createExercise()
			workout.Exercises = append(workout.Exercises, exercise)
		} else {
			break
		}
	}

	createError := storage.Create(workout)

	if createError != nil {
		ui.PrintError(fmt.Sprintf("Failed to create workout: %v", createError))
		return
	}

	ui.PrintSuccess(fmt.Sprintf("Workout '%s' created successfully!", workout.Name))
}

func createExercise() workouts.Exercise {
	ui.PrintInfo("Adding new exercise...")

	name := ui.PromptString("Exercise name")
	sets := ui.PromptInt("Number of sets")
	reps := ui.PromptInt("Number of reps per set")
	weight := ui.PromptInt("Weight(kg)")
	restTime := ui.PromptInt("Rest time between sets (seconds)")
	description := ui.PromptString("Exercise notes (optional)")

	return workouts.Exercise{
		Name:        name,
		Sets:        sets,
		Reps:        reps,
		Weight:      weight,
		RestTime:    restTime,
		Description: description,
	}
}

func listWorkoutsHandler(cmd *cobra.Command, args []string) {
	storage := workouts.NewJSONStorage()

	workoutList, err := storage.GetAll()
	if err != nil {
		ui.PrintError(fmt.Sprintf("Failed to load workouts: %v", err))
		return
	}

	if len(workoutList) == 0 {
		ui.PrintInfo("No workouts found. Create your first workout with 'gain-train workouts add'")
		return
	}

	fmt.Println("\n Your Workouts:")
	fmt.Println("================")

	for i, workout := range workoutList {
		fmt.Printf("\n%d. %s\n", i+1, workout.Name)
		if workout.Description != "" {
			fmt.Printf("   Description: %s\n", workout.Description)
		}
		fmt.Printf("   Exercises: %d\n", len(workout.Exercises))
		fmt.Printf("   Created: %s\n", workout.CreatedAt.Format("2006-01-02 15:04"))

		// Show exercises
		for j, exercise := range workout.Exercises {
			fmt.Printf("     %d. %s - %d sets x %d reps @ %d kg\n",
				j+1, exercise.Name, exercise.Sets, exercise.Reps, exercise.Weight)
		}
	}
}

func updateWorkoutHandler(cmd *cobra.Command, args []string) {
	fmt.Println("Updating a workout...")
	// TODO: Implement update workout logic
}

func removeWorkoutHandler(cmd *cobra.Command, args []string) {
	storage := workouts.NewJSONStorage()

	workoutList, err := storage.GetAll()
	if err != nil {
		ui.PrintError(fmt.Sprintf("Failed to load workouts: %v", err))
		return
	}

	if len(workoutList) == 0 {
		ui.PrintInfo("No workouts to remove.")
		return
	}

	// Show available workouts
	fmt.Println("\nSelect a workout to remove:")
	for i, workout := range workoutList {
		fmt.Printf("%d. %s\n", i+1, workout.Name)
	}

	// Get user selection
	choice := ui.PromptInt("Enter workout number")
	if choice < 1 || choice > len(workoutList) {
		ui.PrintError("Invalid workout number.")
		return
	}

	selectedWorkout := workoutList[choice-1]

	// Confirm deletion
	if !ui.PromptConfirm(fmt.Sprintf("Are you sure you want to delete '%s'?", selectedWorkout.Name)) {
		ui.PrintInfo("Deletion cancelled.")
		return
	}

	// Delete workout
	if err := storage.Delete(selectedWorkout.ID); err != nil {
		ui.PrintError(fmt.Sprintf("Failed to delete workout: %v", err))
		return
	}

	ui.PrintSuccess(fmt.Sprintf("Workout '%s' deleted successfully!", selectedWorkout.Name))
}
