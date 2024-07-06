package domain

type WodSets struct {
	SetNumber        int32 `json:"set_number" binding:"required"`
	BuyIn            int32 `json:"buy_in"`
	BuyOut           int32 `json:"buy_out"`
	EveryMinutes     int32 `json:"every_minutes"`
	RepsToAddByRound int32 `json:"reps_to_add_by_round"`
	RestTime         int32 `json:"rest_time"`
	IsThen           bool  `json:"is_then"`
}

type WodRounds struct {
	SetNumber          int32 `json:"set_number" binding:"required"`
	RoundNumber        int32 `json:"round_number" binding:"required"`
	RepetitionsByRound int32 `json:"repetitions_by_round"`
	Time               int32 `json:"time"`
	RestTime           int32 `json:"rest_time"`
	RemainingTime      bool  `json:"remaining_Time"`
}

type ExerciseInWod struct {
	ID              int32   `json:"id" binding:"required"`
	SetNumber       int32   `json:"set_number" binding:"required"`
	RoundNumber     int32   `json:"round_number" binding:"required"`
	Repetitions     int32   `json:"repetitions" binding:"required"`
	RepetitionsUnit string  `json:"repetitions_unit" binding:"required"`
	Weight          float32 `json:"weight"`
	WeightUnit      string  `json:"weight_unit"`
	Distance        float32 `json:"distance"`
	DistanceUnit    string  `json:"distance_unit"`
}

type CreateWodRequest struct {
	ID        string          `json:"id" binding:"required"`
	Name      string          `json:"name" binding:"required"`
	WodDate   string          `json:"wod_date"`
	WodTypeId int32           `json:"wod_type_id" binding:"required"`
	Sets      []WodSets       `json:"sets" binding:"required"`
	Rounds    []WodRounds     `json:"rounds" binding:"required"`
	Exercises []ExerciseInWod `json:"exercises" binding:"required"`
}

type CriteriaFilters struct {
	field    string `json:"field"`
	operator string `json:"operator"`
	value    string `json:"value"`
}

type CriteriaRequest struct {
	filters []CriteriaFilters `json:"filters" binding:"required"`
	orderBy string            `json:"order_by" binding:"required"`
	order   string            `json:"order" binding:"required"`
}
