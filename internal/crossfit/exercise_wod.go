package crossfit

type ExerciseWod struct {
	ID              string `gorm:"primaryKey"`
	WodId           string
	ExerciseId      uint16
	SetNumber       uint8
	RoundNumber     uint8
	Repetitions     uint16
	RepetitionsUnit string
	Weight          float32
	WeightUnit      string
	Distance        float32
	DistanceUnit    string
}

func NewExerciseWod(id, wodId string, exerciseId uint16, setNumber, roundNumber uint8, reps uint16, repsUnit string, weight float32, weightUnit string, distance float32, distanceUnit string) (ExerciseWod, error) {
	return ExerciseWod{
		ID:              id,
		WodId:           wodId,
		ExerciseId:      exerciseId,
		SetNumber:       setNumber,
		RoundNumber:     roundNumber,
		Repetitions:     reps,
		RepetitionsUnit: repsUnit,
		Weight:          weight,
		WeightUnit:      weightUnit,
		Distance:        distance,
		DistanceUnit:    distanceUnit,
	}, nil

}

type ExerciseWodRepository interface {
	Save(exerciseWod ExerciseWod) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=../platform/storage/storagemocks --name=ExerciseWodRepository
