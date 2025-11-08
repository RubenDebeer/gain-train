package program

import (
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
		filePath: filepath.Join(dataDir, "program.json"),
	}
}

func (js *JSONStorage) Create(program *Program) error {

	// Create ID
	// Update Date Values
	// Save all the shit
	// Save workouts
	// Save Excercises

	program.ID = uuid.New().String()
	program.CreatedAt = time.Now()
	program.UpdatedAt = time.Now()

	// Get all Workouts
	// Get all of the Programs

	return nil
}

// Save helper
// Load Helper
