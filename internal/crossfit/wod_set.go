package crossfit

type WodSet struct {
	ID               string `gorm:"primaryKey"`
	WodId            string
	SetNumber        uint8
	BuyIn            uint16
	BuyOut           uint16
	EveryMinutes     uint32
	RepsToAddByRound uint16
	RestTime         uint32
	IsThen           bool
}

func NewWodSet(id, wodId string, setNumber uint8, buyIn, buyOut uint16, everyMinutes uint32, repsToAddByRound uint16, restTime uint32, isThen bool) (WodSet, error) {

	return WodSet{
		ID:               id,
		WodId:            wodId,
		SetNumber:        setNumber,
		BuyIn:            buyIn,
		BuyOut:           buyOut,
		EveryMinutes:     everyMinutes,
		RepsToAddByRound: repsToAddByRound,
		RestTime:         restTime,
		IsThen:           isThen,
	}, nil
}

type WodSetRepository interface {
	Save(w WodSet) error
	// Get(id string) (*WodSet, error)
}
