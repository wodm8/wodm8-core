package mysql

const (
	sqlWodRoundTable = "WOD_ROUND"
)

type sqlWodRound struct {
	ID                 string `db:"id"`
	WodId              string `db:"wod_id"`
	SetNumber          int32  `db:"set_number"`
	RoundNumber        int32  `db:"round_number"`
	RepetitionsByRound int32  `db:"repetitions_by_round"`
	Time               int32  `db:"time"`
	RestTime           int32  `db:"rest_time"`
	RemainingRest      bool   `db:"remaining_time"`
}
