package crossfit

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var ErrInvalidWodId = errors.New("invalid wod id")

type WodId struct {
	value string
}

func NewWodId(value string) (WodId, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return WodId{}, fmt.Errorf("%w: %s", ErrInvalidWodId, value)
	}
	return WodId{
		value: v.String(),
	}, nil
}

func (id WodId) String() string {
	return id.value
}

var ErrInvalidWodName = errors.New("invalid wod name")

type WodName struct {
	value string
}

func NewWodName(value string) (WodName, error) {
	if value == "" {
		return WodName{}, fmt.Errorf("%w: %s", ErrInvalidWodName, value)
	}
	return WodName{
		value: value,
	}, nil
}

func (name WodName) String() string {
	return name.value
}

var ErrInvalidWodRounds = errors.New("invalid wod rounds")

type WodRounds struct {
	value int32
}

func NewWodRounds(value int32) (WodRounds, error) {
	if value == 0 {
		return WodRounds{}, fmt.Errorf("%w: %d", ErrInvalidWodRounds, value)
	}
	return WodRounds{
		value: value,
	}, nil
}

func (rounds WodRounds) Int() int32 {
	return rounds.value
}

var ErrInvalidWodSections = errors.New("invalid number of sections in wod")

type WodNumberSections struct {
	value int32
}

func NewWodNumberSections(value int32) (WodNumberSections, error) {
	if value == 0 {
		return WodNumberSections{}, fmt.Errorf("%w: %d", ErrInvalidWodSections, value)
	}
	return WodNumberSections{
		value: value,
	}, nil
}

func (sections WodNumberSections) Int() int32 {
	return sections.value
}

var ErrInvalidTimerType = errors.New("invalid timer type id")

type WodTimerTypeID struct {
	value int32
}

func NewWodTimerTypeID(value int32) (WodTimerTypeID, error) {
	if value == 0 {
		return WodTimerTypeID{}, fmt.Errorf("%w: %d", ErrInvalidTimerType, value)
	}
	return WodTimerTypeID{
		value: value,
	}, nil
}

func (timerType WodTimerTypeID) Int() int32 {
	return timerType.value
}

type Wod struct {
	id             WodId
	name           WodName
	rounds         WodRounds
	numberSections WodNumberSections
	timerTypeId    WodTimerTypeID
}

func NewWod(id, wodName string, rounds, numberSections, timerTypeId int32) (Wod, error) {
	idVO, err := NewWodId(id)
	if err != nil {
		return Wod{}, err
	}

	nameVO, err := NewWodName(wodName)
	if err != nil {
		return Wod{}, err
	}

	roundsVO, err := NewWodRounds(rounds)
	if err != nil {
		return Wod{}, err
	}

	nSectionsVO, err := NewWodNumberSections(numberSections)
	if err != nil {
		return Wod{}, err
	}

	tTypeIdVO, err := NewWodTimerTypeID(timerTypeId)
	if err != nil {
		return Wod{}, err
	}

	return Wod{
		id:             idVO,
		name:           nameVO,
		rounds:         roundsVO,
		numberSections: nSectionsVO,
		timerTypeId:    tTypeIdVO,
	}, err
}

type WodRepository interface {
	Save(ctx context.Context, wod Wod) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=ExerciseRepository

func (w Wod) ID() WodId {
	return w.id
}

func (w Wod) Name() WodName {
	return w.name
}

func (w Wod) Rounds() WodRounds {
	return w.rounds
}

func (w Wod) NumberSections() WodNumberSections {
	return w.numberSections
}

func (w Wod) TimerType() WodTimerTypeID {
	return w.timerTypeId
}
