package crossfit

import (
	"context"
	"errors"
	"fmt"
	"github.com/wodm8/wodm8-core/internal/domain"

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

var ErrInvalidWodDate = errors.New("invalid wod date")

type WodDate struct {
	value string
}

func NewWodDate(value string) (WodDate, error) {
	if value == "" {
		return WodDate{}, fmt.Errorf("%w: %s", ErrInvalidWodDate, value)
	}
	return WodDate{
		value: value,
	}, nil
}

func (wd WodDate) String() string {
	return wd.value
}

var ErrInvalidWodType = errors.New("invalid wod type id")

type WodType struct {
	value int32
}

func NewWodType(value int32) (WodType, error) {
	if value == 0 {
		return WodType{}, fmt.Errorf("%w: %d", ErrInvalidWodType, value)
	}
	return WodType{
		value: value,
	}, nil
}

func (wt WodType) Int() int32 {
	return wt.value
}

type Wod struct {
	id        WodId
	name      WodName
	wodDate   WodDate
	wodTypeId WodType
}

func NewWod(id, wodName, date string, wodTypeId int32) (Wod, error) {
	idVO, err := NewWodId(id)
	if err != nil {
		return Wod{}, err
	}

	nameVO, err := NewWodName(wodName)
	if err != nil {
		return Wod{}, err
	}

	dateVO, err := NewWodDate(date)
	if err != nil {
		return Wod{}, err
	}

	wTypeVO, err := NewWodType(wodTypeId)
	if err != nil {
		return Wod{}, err
	}

	return Wod{
		id:        idVO,
		name:      nameVO,
		wodDate:   dateVO,
		wodTypeId: wTypeVO,
	}, err
}

type WodRepository interface {
	Save(ctx context.Context, wod Wod) error
	Get(ctx context.Context, id string) ([]domain.WodResult, error)
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=ExerciseRepository

func (w Wod) ID() WodId {
	return w.id
}

func (w Wod) Name() WodName {
	return w.name
}

func (w Wod) Date() WodDate {
	return w.wodDate
}

func (w Wod) TypeID() WodType {
	return w.wodTypeId
}
