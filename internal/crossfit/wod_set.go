package crossfit

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var ErrInvalidSetId = errors.New("invalid set id")

type SetId struct {
	value string
}

func NewSetId(value string) (SetId, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return SetId{}, fmt.Errorf("%w: %s", ErrInvalidSetId, value)
	}
	return SetId{
		value: v.String(),
	}, nil
}

func (id SetId) String() string {
	return id.value
}

type SetWodId struct {
	value string
}

func NewSetWodId(value string) (SetWodId, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return SetWodId{}, fmt.Errorf("%w: %s", ErrInvalidWodId, value)
	}
	return SetWodId{
		value: v.String(),
	}, nil
}

func (id SetWodId) String() string {
	return id.value
}

var ErrInvalidSetNumber = errors.New("invalid set number")

type SetNumber struct {
	value int32
}

func NewSetNumber(value int32) (SetNumber, error) {
	if value == 0 {
		return SetNumber{}, fmt.Errorf("%w: %d", ErrInvalidSetNumber, value)
	}
	return SetNumber{
		value: value,
	}, nil
}

func (s SetNumber) Int() int32 {
	return s.value
}

type SetBuyIn struct {
	value int32
}

func NewSetBuyIn(value int32) (SetBuyIn, error) {
	return SetBuyIn{
		value: value,
	}, nil
}

func (sbi SetBuyIn) Int() int32 {
	return sbi.value
}

type SetBuyOut struct {
	value int32
}

func NewSetBuyOut(value int32) (SetBuyOut, error) {
	return SetBuyOut{
		value: value,
	}, nil
}

func (sbo SetBuyOut) Int() int32 {
	return sbo.value
}

type SetEveryMinutes struct {
	value int32
}

func NewSetEveryMinutes(value int32) (SetEveryMinutes, error) {
	return SetEveryMinutes{
		value: value,
	}, nil
}

func (srt SetEveryMinutes) Int() int32 {
	return srt.value
}

type SetRepsToAddByRound struct {
	value int32
}

func NewSetRepsToAddByRound(value int32) (SetRepsToAddByRound, error) {
	return SetRepsToAddByRound{
		value: value,
	}, nil
}

func (srt SetRepsToAddByRound) Int() int32 {
	return srt.value
}

type SetRestTime struct {
	value int32
}

func NewSetRestTime(value int32) (SetRestTime, error) {
	return SetRestTime{
		value: value,
	}, nil
}

func (srt SetRestTime) Int() int32 {
	return srt.value
}

type SetIsThen struct {
	value bool
}

func NewSetIsThen(value bool) (SetIsThen, error) {
	return SetIsThen{
		value: value,
	}, nil
}

func (sit SetIsThen) Bool() bool {
	return sit.value
}

type WodSet struct {
	id               SetId
	wodId            SetWodId
	setNumber        SetNumber
	buyIn            SetBuyIn
	buyOut           SetBuyOut
	everyMinutes     SetEveryMinutes
	repsToAddByRound SetRepsToAddByRound
	restTime         SetRestTime
	isThen           SetIsThen
}

func NewWodSet(id, wodId string, setNumber, buyIn, buyOut, everyMinutes, repsToAddByR, restTime int32, isThen bool) (WodSet, error) {
	idVO, err := NewSetId(id)
	if err != nil {
		return WodSet{}, err
	}

	setWodIdVO, err := NewSetWodId(wodId)
	if err != nil {
		return WodSet{}, err
	}

	setNumberVO, err := NewSetNumber(setNumber)
	if err != nil {
		return WodSet{}, err
	}

	buyInVO, err := NewSetBuyIn(buyIn)
	if err != nil {
		return WodSet{}, err
	}

	buyOutVO, err := NewSetBuyOut(buyOut)
	if err != nil {
		return WodSet{}, err
	}

	everyMinutesVO, err := NewSetEveryMinutes(everyMinutes)
	if err != nil {
		return WodSet{}, err
	}

	repsToAddVO, err := NewSetRepsToAddByRound(repsToAddByR)
	if err != nil {
		return WodSet{}, err
	}

	restTimeVO, err := NewSetRestTime(restTime)
	if err != nil {
		return WodSet{}, err
	}

	isThenVO, err := NewSetIsThen(isThen)
	if err != nil {
		return WodSet{}, err
	}

	return WodSet{
		id:               idVO,
		wodId:            setWodIdVO,
		setNumber:        setNumberVO,
		buyIn:            buyInVO,
		buyOut:           buyOutVO,
		everyMinutes:     everyMinutesVO,
		repsToAddByRound: repsToAddVO,
		restTime:         restTimeVO,
		isThen:           isThenVO,
	}, nil
}

func (w WodSet) ID() SetId {
	return w.id
}

func (w WodSet) WodId() SetWodId {
	return w.wodId
}

func (w WodSet) SetNumber() SetNumber {
	return w.setNumber
}
func (w WodSet) SetBuyIn() SetBuyIn                       { return w.buyIn }
func (w WodSet) SetBuyOut() SetBuyOut                     { return w.buyOut }
func (w WodSet) SetEveryMinutes() SetEveryMinutes         { return w.everyMinutes }
func (w WodSet) SetRepsToAddByRound() SetRepsToAddByRound { return w.repsToAddByRound }
func (w WodSet) SetRestTime() SetRestTime                 { return w.restTime }
func (w WodSet) SetIsThen() SetIsThen                     { return w.isThen }

type WodSetRepository interface {
	Save(ctx context.Context, w WodSet) error
	// Get(id string) (*WodSet, error)
}
