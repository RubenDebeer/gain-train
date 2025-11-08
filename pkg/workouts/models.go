package workouts

import "time"

// Workout represents a complete workout session
type Workout struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Exercises   []Exercise `json:"exercises"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// Exercise represents a single exercise within a workout
type Exercise struct {
	Name        string `json:"name"`
	Sets        int    `json:"sets"`
	Reps        int    `json:"reps"`
	Weight      int    `json:"weight"`    // in lbs or kg
	Duration    int    `json:"duration"`  // in seconds for time-based exercises
	RestTime    int    `json:"rest_time"` // in seconds
	Description string `json:"description"`
}

// WorkoutService defines the interface for workout operations
type WorkoutService interface {
	Create(workout *Workout) error
	GetAll() ([]Workout, error)
	GetByID(id string) (*Workout, error)
	Update(workout *Workout) error
	Delete(id string) error
}
