package crossfit

type Exercise struct {
	Name string
}

// NewExercise creates a new exercise
func NewExercise(name string) (Exercise, error) {
	return Exercise{
		Name: name,
	}, nil
}

type ExerciseRepository interface {
	Save(exercise Exercise) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=../platform/storage/storagemocks --name=ExerciseRepository
