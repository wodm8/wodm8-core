package domain

type WodResult struct {
	WodId                  string  `json:"id"`
	WodName                string  `json:"name"`
	WodType                string  `json:"wod_type"`
	TypeDescription        string  `json:"type_description"`
	SetNumber              int32   `json:"set_number"`
	BuyIn                  int32   `json:"buy_in"`
	BuyOut                 int32   `json:"buy_out"`
	EveryMinutes           int32   `json:"every_minutes"`
	RepsToAddByRound       int32   `json:"reps_to_add_by_round"`
	SetRestTime            int32   `json:"set_rest_time"`
	IsThen                 bool    `json:"is_then"`
	RoundNumber            int32   `json:"round_number"`
	RepetitionsByRound     int32   `json:"repetitions_by_round"`
	RoundTime              int32   `json:"round_time"`
	RestRemainingRoundTime bool    `json:"rest_remaining_round_time"`
	ExerciseName           string  `json:"exercise_name"`
	Repetitions            int32   `json:"repetitions"`
	RepetitionsUnit        string  `json:"repetitions_unit"`
	Weight                 float32 `json:"weight"`
	WeightUnit             string  `json:"weight_unit"`
	Distance               float32 `json:"distance"`
	DistanceUnit           string  `json:"distance_unit"`
}
