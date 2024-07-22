package crossfit

type WodRound struct {
	ID                 string `gorm:"primaryKey"`
	WodId              string
	SetNumber          uint8
	RoundNumber        uint8
	RepetitionsByRound uint16
	Time               uint32
	RestTime           uint32
	RemainingTime      bool
}

func NewWodRound(id, wodId string, setNumber, roundNumber uint8, repetitionsByRound uint16, time, restTime uint32, remainingTime bool) (WodRound, error) {

	return WodRound{
		ID:                 id,
		WodId:              wodId,
		SetNumber:          setNumber,
		RoundNumber:        roundNumber,
		RepetitionsByRound: repetitionsByRound,
		Time:               time,
		RestTime:           restTime,
		RemainingTime:      remainingTime,
	}, nil
}

type WodRoundRepository interface {
	Save(wr WodRound) error
}
