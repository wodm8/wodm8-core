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
	value string
}

func NewExWodExerciseID(value string) (ExWodExerciseID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return ExWodExerciseID{}, fmt.Errorf("%w: %s", ErrInvalidExerciseWodExerciseId, value)
	}
	return ExWodExerciseID{
		value: v.String(),
	}, nil
}

func (id ExWodExerciseID) String() string {
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

var ErrInvalidExerciseWodWeight = errors.New("invalid weight")

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

var ErrInvalidExerciseWodWeightU = errors.New("invalid weight unit")

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

var ErrInvalidExerciseWodSection = errors.New("invalid section")

type ExWodSection struct {
	value int32
}

func NewExWodSection(value int32) (ExWodSection, error) {
	if value == 0 {
		return ExWodSection{}, fmt.Errorf("%w: %d", ErrInvalidExerciseWodSection, value)
	}
	return ExWodSection{
		value: value,
	}, nil
}

func (id ExWodSection) Int() int32 {
	return id.value
}

var ErrInvalidExerciseWodTimer = errors.New("invalid timer type")

type ExWodSectionTimerType struct {
	value int32
}

func NewExWodSectionTimerType(value int32) (ExWodSectionTimerType, error) {
	if value == 0 {
		return ExWodSectionTimerType{}, fmt.Errorf("%w: %d", ErrInvalidExerciseWodTimer, value)
	}
	return ExWodSectionTimerType{
		value: value,
	}, nil
}

func (id ExWodSectionTimerType) Int() int32 {
	return id.value
}

var ErrInvalidExerciseWodCap = errors.New("invalid Cap")

type ExWodSectionCap struct {
	value int32
}

func NewExWodSectionCap(value int32) (ExWodSectionCap, error) {
	return ExWodSectionCap{
		value: value,
	}, nil
}

func (id ExWodSectionCap) Int() int32 {
	return id.value
}

type ExerciseWod struct {
	id               ExWodID
	wodId            ExWodReferID
	exerciseId       ExWodExerciseID
	repetitions      ExWodReps
	weight           ExWodWeight
	weightUnit       ExWodWeightUnit
	wodSection       ExWodSection
	sectionTimerType ExWodSectionTimerType
	sectionCap       ExWodSectionCap
}

func NewExerciseWod(id, wodId, exerciseId string, reps int32, weight float32, weightUnit string, section, timer, cap int32) (ExerciseWod, error) {
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

	repsVO, err := NewExWodRepetitions(reps)
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

	sectionVO, err := NewExWodSection(section)
	if err != nil {
		return ExerciseWod{}, err
	}

	timerVO, err := NewExWodSectionTimerType(timer)
	if err != nil {
		return ExerciseWod{}, err
	}

	capVO, err := NewExWodSectionCap(cap)
	if err != nil {
		return ExerciseWod{}, err
	}

	return ExerciseWod{
		id:               idVO,
		wodId:            wodIdVO,
		exerciseId:       exerciseIdVO,
		repetitions:      repsVO,
		weight:           weightVO,
		weightUnit:       weightUnitVO,
		wodSection:       sectionVO,
		sectionTimerType: timerVO,
		sectionCap:       capVO,
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

func (e ExerciseWod) Repetitions() ExWodReps {
	return e.repetitions
}

func (e ExerciseWod) Weight() ExWodWeight {
	return e.weight
}

func (e ExerciseWod) WeightUnit() ExWodWeightUnit {
	return e.weightUnit
}

func (e ExerciseWod) Section() ExWodSection {
	return e.wodSection
}

func (e ExerciseWod) TimerType() ExWodSectionTimerType {
	return e.sectionTimerType
}

func (e ExerciseWod) Cap() ExWodSectionCap {
	return e.sectionCap
}
