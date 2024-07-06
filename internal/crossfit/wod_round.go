package crossfit

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

var ErrWRID = errors.New("WRID is not valid")

type WRID struct {
	value string
}

func NewWRID(value string) (WRID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return WRID{}, fmt.Errorf("%w: %s", ErrWRID, value)
	}
	return WRID{value: v.String()}, nil
}

func (wr WRID) String() string {
	return wr.value
}

var ErrWRWodID = errors.New("WR wod id is not valid")

type WRWodID struct {
	value string
}

func NewWRWodID(value string) (WRWodID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return WRWodID{}, fmt.Errorf("%w: %s", ErrWRWodID, value)
	}
	return WRWodID{value: v.String()}, nil
}

func (wrw WRWodID) String() string {
	return wrw.value
}

var ErrWRSetNumber = errors.New("WR set number is not valid")

type WRSetNumber struct {
	value int32
}

func NewWRSetNumber(value int32) (WRSetNumber, error) {
	if value <= 0 {
		return WRSetNumber{}, fmt.Errorf("%w: %d", ErrWRSetNumber, value)
	}
	return WRSetNumber{value: value}, nil
}

func (wrsn WRSetNumber) Int() int32 {
	return wrsn.value
}

var ErrWRRounds = errors.New("WR rounds is not valid")

type WRRoundNumber struct {
	value int32
}

func NewWRRoundNumber(value int32) (WRRoundNumber, error) {
	if value <= 0 {
		return WRRoundNumber{}, fmt.Errorf("%w: %d", ErrWRRounds, value)
	}
	return WRRoundNumber{value: value}, nil
}

func (wrr WRRoundNumber) Int() int32 {
	return wrr.value
}

type WRRepsNumber struct {
	value int32
}

func NewWRRepsNumber(value int32) (WRRepsNumber, error) {
	return WRRepsNumber{value: value}, nil
}

func (wrn WRRepsNumber) Int() int32 {
	return wrn.value
}

type WRTime struct {
	value int32
}

func NewWRTime(value int32) (WRTime, error) {
	return WRTime{value: value}, nil
}

func (wrt WRTime) Int() int32 {
	return wrt.value
}

type WRRestTime struct {
	value int32
}

func NewWRRestTime(value int32) (WRRestTime, error) {
	return WRRestTime{value: value}, nil
}

func (wrrt WRRestTime) Int() int32 {
	return wrrt.value
}

type WRRemainingTime struct {
	value bool
}

func NewWRRemainingTime(value bool) (WRRemainingTime, error) {
	return WRRemainingTime{value: value}, nil
}

func (wrt WRRemainingTime) Bool() bool {
	return wrt.value
}

type WodRound struct {
	id                 WRID
	wodId              WRWodID
	setNumber          WRSetNumber
	roundNumber        WRRoundNumber
	repetitionsByRound WRRepsNumber
	time               WRTime
	restTime           WRRestTime
	remainingTime      WRRemainingTime
}

func NewWodRound(id, WodId string, setNumber, roundNumber, repetitionsByRound, time, restTime int32, remainingTime bool) (WodRound, error) {
	idVO, err := NewWRID(id)
	if err != nil {
		return WodRound{}, err
	}

	wodIdVO, err := NewWRWodID(WodId)
	if err != nil {
		return WodRound{}, err
	}

	setNumberVO, err := NewWRSetNumber(setNumber)
	if err != nil {
		return WodRound{}, err
	}

	roundNumberVO, err := NewWRRoundNumber(roundNumber)
	if err != nil {
		return WodRound{}, err
	}

	repetitionsByRoundVO, err := NewWRRepsNumber(repetitionsByRound)
	if err != nil {
		return WodRound{}, err
	}

	timeVO, err := NewWRTime(time)
	if err != nil {
		return WodRound{}, err
	}

	restTimeVO, err := NewWRRestTime(restTime)
	if err != nil {
		return WodRound{}, err
	}

	remainingTimeVO, err := NewWRRemainingTime(remainingTime)
	if err != nil {
		return WodRound{}, err
	}

	return WodRound{
		id:                 idVO,
		wodId:              wodIdVO,
		setNumber:          setNumberVO,
		roundNumber:        roundNumberVO,
		repetitionsByRound: repetitionsByRoundVO,
		time:               timeVO,
		restTime:           restTimeVO,
		remainingTime:      remainingTimeVO,
	}, nil
}

func (wr WodRound) ID() WRID {
	return wr.id
}

func (wr WodRound) WODId() WRWodID {
	return wr.wodId
}

func (wr WodRound) SetNumber() WRSetNumber {
	return wr.setNumber
}

func (wr WodRound) RoundNumber() WRRoundNumber {
	return wr.roundNumber
}

func (wr WodRound) RepetitionsByRound() WRRepsNumber {
	return wr.repetitionsByRound
}

func (wr WodRound) Time() WRTime {
	return wr.time
}

func (wr WodRound) RestTime() WRRestTime {
	return wr.restTime
}

func (wr WodRound) RemainingTime() WRRemainingTime {
	return wr.remainingTime
}

type WodRoundRepository interface {
	Save(ctx context.Context, wr WodRound) error
}
