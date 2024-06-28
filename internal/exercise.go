package crossfit

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var ErrInvalidExerciseID = errors.New("invalid Exercise ID")

type ExerciseId struct {
	value string
}

func NewExerciseId(value string) (ExerciseId, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return ExerciseId{}, fmt.Errorf("%w: %s", ErrInvalidExerciseID, value)
	}
	return ExerciseId{
		value: v.String(),
	}, nil
}

func (id ExerciseId) String() string {
	return id.value
}

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
	id   ExerciseId
	name ExerciseName
}

// NewExercise creates a new exercise
func NewExercise(id, name string) (Exercise, error) {
	idVO, err := NewExerciseId(id)
	if err != nil {
		return Exercise{}, err
	}

	NameVO, err := NewExerciseName(name)
	if err != nil {
		return Exercise{}, err
	}

	return Exercise{
		id:   idVO,
		name: NameVO,
	}, nil
}

type ExerciseRepository interface {
	Save(ctx context.Context, exercise Exercise) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=ExerciseRepository

// ID return the exercise unique identifier
func (e Exercise) ID() ExerciseId {
	return e.id
}

// Name return the exercise name
func (e Exercise) Name() ExerciseName {
	return e.name
}
