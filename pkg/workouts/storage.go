package workouts

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type JSONStorage struct {
	filePath string
}

func NewJSONStorage() *JSONStorage {
	homeDir, _ := os.UserHomeDir()
	dataDir := filepath.Join(homeDir, ".gain-train")
	os.MkdirAll(dataDir, 0755)

	return &JSONStorage{
		filePath: filepath.Join(dataDir, "workouts.json"),
	}
}

func (js *JSONStorage) Create(workout *Workout) error {
	// Generate ID and timestamps
	workout.ID = uuid.New().String()
	workout.CreatedAt = time.Now()
	workout.UpdatedAt = time.Now()

	workouts, err := js.loadWorkouts()
	if err != nil {
		return err
	}

	workouts = append(workouts, *workout)
	return js.saveWorkouts(workouts)
}

func (js *JSONStorage) GetAll() ([]Workout, error) {
	return js.loadWorkouts()
}

func (js *JSONStorage) GetByID(id string) (*Workout, error) {
	workouts, err := js.loadWorkouts()
	if err != nil {
		return nil, err
	}

	for _, workout := range workouts {
		if workout.ID == id {
			return &workout, nil
		}
	}

	return nil, fmt.Errorf("workout with ID %s not found", id)
}

func (js *JSONStorage) Update(workout *Workout) error {
	workouts, err := js.loadWorkouts()
	if err != nil {
		return err
	}

	for i, w := range workouts {
		if w.ID == workout.ID {
			workout.UpdatedAt = time.Now()
			workouts[i] = *workout
			return js.saveWorkouts(workouts)
		}
	}

	return fmt.Errorf("workout with ID %s not found", workout.ID)
}

func (js *JSONStorage) Delete(id string) error {
	workouts, err := js.loadWorkouts()
	if err != nil {
		return err
	}

	for i, workout := range workouts {
		if workout.ID == id {
			workouts = append(workouts[:i], workouts[i+1:]...)
			return js.saveWorkouts(workouts)
		}
	}

	return fmt.Errorf("workout with ID %s not found", id)
}

func (js *JSONStorage) loadWorkouts() ([]Workout, error) {
	var workouts []Workout

	if _, err := os.Stat(js.filePath); os.IsNotExist(err) {
		return workouts, nil
	}

	data, err := os.ReadFile(js.filePath)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return workouts, nil
	}

	err = json.Unmarshal(data, &workouts)
	return workouts, err
}

func (js *JSONStorage) saveWorkouts(workouts []Workout) error {
	data, err := json.MarshalIndent(workouts, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(js.filePath, data, 0644)
}
