package crossfit

import (
	"context"
	"errors"
)

var ErrEmptyExerciseName = errors.New("invalid Exercise ID")

type ExerciseName struct {
	value string
}

func NewExerciseName(value string) (ExerciseName, error) {
	if value == "" {
		return ExerciseName{}, ErrEmptyExerciseName
	}
	return ExerciseName{
		value: value,
	}, nil
}

func (name ExerciseName) String() string {
	return name.value
}

type Exercise struct {
	name ExerciseName
}

// NewExercise creates a new exercise
func NewExercise(name string) (Exercise, error) {
	NameVO, err := NewExerciseName(name)
	if err != nil {
		return Exercise{}, err
	}

	return Exercise{
		name: NameVO,
	}, nil
}

type ExerciseRepository interface {
	Save(ctx context.Context, exercise Exercise) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=ExerciseRepository

// Name return the exercise name
func (e Exercise) Name() ExerciseName {
	return e.name
}
