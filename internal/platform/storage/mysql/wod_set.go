package mysql

const (
	sqlWodSetTable = "WOD_SET"
)

type sqlWodSet struct {
	ID               string `db:"id"`
	WodId            string `db:"wod_id"`
	SetNumber        int32  `db:"set_number"`
	BuyIn            int32  `db:"buy_in"`
	BuyOut           int32  `db:"buy_out"`
	EveryMinutes     int32  `db:"every_minutes"`
	RepsToAddByRound int32  `db:"reps_to_add_by_round"`
	RestTime         int32  `db:"rest_time"`
	IsThen           bool   `db:"is_then"`
}
