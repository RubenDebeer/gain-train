package program

import (
	"gain-train-cli/pkg/workouts"
	"time"
)

type Program struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Workouts    []workouts.Workout `json:"workouts"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

type ProgramService interface {
	Create(program *Program) error
	GetAll() ([]Program, error)
	GetByID() (Program, error)
	Update(id string) error
	Delete(Program *Program) error
}
