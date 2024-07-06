package crossfit

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var ErrInvalidExerciseWodId = errors.New("invalid ExerciseWod ID")

type ExWodID struct {
	value string
}

func NewExWodId(value string) (ExWodID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return ExWodID{}, fmt.Errorf("%w: %s", ErrInvalidExerciseWodId, value)
	}
	return ExWodID{
		value: v.String(),
	}, nil
}

func (id ExWodID) String() string {
	return id.value
}

var ErrExWodSetNumber = errors.New("ExW set number is not valid")

type ExWodSetNumber struct {
	value int32
}

func NewExWodSetNumber(value int32) (ExWodSetNumber, error) {
	if value <= 0 {
		return ExWodSetNumber{}, fmt.Errorf("%w: %d", ErrExWodSetNumber, value)
	}
	return ExWodSetNumber{value: value}, nil
}

func (exwsn ExWodSetNumber) Int() int32 {
	return exwsn.value
}

var ErrInvalidExerciseWodReferId = errors.New("invalid exercise ID")

type ExWodReferID struct {
	value string
}

func NewExWodReferID(value string) (ExWodReferID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return ExWodReferID{}, fmt.Errorf("%w: %s", ErrInvalidExerciseWodReferId, value)
	}
	return ExWodReferID{
		value: v.String(),
	}, nil
}

func (id ExWodReferID) String() string {
	return id.value
}

var ErrInvalidExerciseWodExerciseId = errors.New("invalid Exercise Refer ID")

type ExWodExerciseID struct {
	value int32
}

func NewExWodExerciseID(value int32) (ExWodExerciseID, error) {
	if value == 0 {
		return ExWodExerciseID{}, fmt.Errorf("%w: %d", ErrInvalidExerciseWodExerciseId, value)
	}
	return ExWodExerciseID{
		value: value,
	}, nil
}

func (eid ExWodExerciseID) Int() int32 {
	return eid.value
}

var ErrInvalidRoundNumber = errors.New("invalid round number")

type ExWodRoundNumber struct {
	value int32
}

func NewExWodRoundNumber(value int32) (ExWodRoundNumber, error) {
	if value == 0 {
		return ExWodRoundNumber{}, fmt.Errorf("%w: %d", ErrInvalidRoundNumber, value)
	}
	return ExWodRoundNumber{
		value: value,
	}, nil
}

func (id ExWodRoundNumber) Int() int32 {
	return id.value
}

var ErrInvalidExerciseWodReps = errors.New("invalid reps")

type ExWodReps struct {
	value int32
}

func NewExWodRepetitions(value int32) (ExWodReps, error) {
	if value == 0 {
		return ExWodReps{}, fmt.Errorf("%w: %d", ErrInvalidExerciseWodReps, value)
	}
	return ExWodReps{
		value: value,
	}, nil
}

func (id ExWodReps) Int() int32 {
	return id.value
}

type ExWodRepsUnit struct {
	value string
}

func NewExWodRepsUnit(value string) (ExWodRepsUnit, error) {
	return ExWodRepsUnit{
		value: value,
	}, nil
}

func (id ExWodRepsUnit) String() string {
	return id.value
}

type ExWodWeight struct {
	value float32
}

func NewExWodWeight(value float32) (ExWodWeight, error) {
	return ExWodWeight{
		value: value,
	}, nil
}

func (id ExWodWeight) Float() float32 {
	return id.value
}

type ExWodWeightUnit struct {
	value string
}

func NewExWodWeightUnit(value string) (ExWodWeightUnit, error) {
	return ExWodWeightUnit{
		value: value,
	}, nil
}

func (id ExWodWeightUnit) String() string {
	return id.value
}

type ExWodDistance struct {
	value float32
}

func NewExWodDistance(value float32) (ExWodDistance, error) {
	return ExWodDistance{
		value: value,
	}, nil
}

func (id ExWodDistance) Float() float32 {
	return id.value
}

type ExWodDistanceUnit struct {
	value string
}

func NewExWodDistanceUnit(value string) (ExWodDistanceUnit, error) {
	return ExWodDistanceUnit{
		value: value,
	}, nil
}

func (id ExWodDistanceUnit) String() string {
	return id.value
}

type ExerciseWod struct {
	id              ExWodID
	wodId           ExWodReferID
	exerciseId      ExWodExerciseID
	setNumber       ExWodSetNumber
	roundNumber     ExWodRoundNumber
	repetitions     ExWodReps
	repetitionsUnit ExWodRepsUnit
	weight          ExWodWeight
	weightUnit      ExWodWeightUnit
	distance        ExWodDistance
	distanceUnit    ExWodDistanceUnit
}

func NewExerciseWod(id, wodId string, exerciseId, setNumber, roundNumber, reps int32, repsUnit string, weight float32, weightUnit string, distance float32, distanceUnit string) (ExerciseWod, error) {
	idVO, err := NewExWodId(id)
	if err != nil {
		return ExerciseWod{}, err
	}

	wodIdVO, err := NewExWodReferID(wodId)
	if err != nil {
		return ExerciseWod{}, err
	}

	exerciseIdVO, err := NewExWodExerciseID(exerciseId)
	if err != nil {
		return ExerciseWod{}, err
	}

	setNumberVO, err := NewExWodSetNumber(setNumber)
	if err != nil {
		return ExerciseWod{}, err
	}

	roundNumberVO, err := NewExWodRoundNumber(roundNumber)
	if err != nil {
		return ExerciseWod{}, err
	}

	repsVO, err := NewExWodRepetitions(reps)
	if err != nil {
		return ExerciseWod{}, err
	}

	repsUnitVO, err := NewExWodRepsUnit(repsUnit)
	if err != nil {
		return ExerciseWod{}, err
	}

	weightVO, err := NewExWodWeight(weight)
	if err != nil {
		return ExerciseWod{}, err
	}

	weightUnitVO, err := NewExWodWeightUnit(weightUnit)
	if err != nil {
		return ExerciseWod{}, err
	}

	distanceVO, err := NewExWodDistance(distance)
	if err != nil {
		return ExerciseWod{}, err
	}

	distanceUnitVO, err := NewExWodDistanceUnit(distanceUnit)
	if err != nil {
		return ExerciseWod{}, err
	}

	return ExerciseWod{
		id:              idVO,
		wodId:           wodIdVO,
		exerciseId:      exerciseIdVO,
		setNumber:       setNumberVO,
		roundNumber:     roundNumberVO,
		repetitions:     repsVO,
		repetitionsUnit: repsUnitVO,
		weight:          weightVO,
		weightUnit:      weightUnitVO,
		distance:        distanceVO,
		distanceUnit:    distanceUnitVO,
	}, nil

}

type ExerciseWodRepository interface {
	Save(ctx context.Context, exerciseWod ExerciseWod) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=ExerciseWodRepository

func (e ExerciseWod) ID() ExWodID {
	return e.id
}

func (e ExerciseWod) WodID() ExWodReferID {
	return e.wodId
}

func (e ExerciseWod) ExerciseID() ExWodExerciseID {
	return e.exerciseId
}

func (e ExerciseWod) SetNumber() ExWodSetNumber {
	return e.setNumber
}

func (e ExerciseWod) RoundNumber() ExWodRoundNumber {
	return e.roundNumber
}

func (e ExerciseWod) Repetitions() ExWodReps {
	return e.repetitions
}

func (e ExerciseWod) RepetitionsUnit() ExWodRepsUnit { return e.repetitionsUnit }

func (e ExerciseWod) Weight() ExWodWeight {
	return e.weight
}

func (e ExerciseWod) WeightUnit() ExWodWeightUnit {
	return e.weightUnit
}

func (e ExerciseWod) Distance() ExWodDistance {
	return e.distance
}

func (e ExerciseWod) DistanceUnit() ExWodDistanceUnit {
	return e.distanceUnit
}
