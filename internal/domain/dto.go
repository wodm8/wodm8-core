package domain

type WodSets struct {
	SetNumber        uint8  `json:"set_number" binding:"required"`
	BuyIn            uint16 `json:"buy_in"`
	BuyOut           uint16 `json:"buy_out"`
	EveryMinutes     uint32 `json:"every_minutes"`
	RepsToAddByRound uint16 `json:"reps_to_add_by_round"`
	RestTime         uint32 `json:"rest_time"`
	IsThen           bool   `json:"is_then"`
}

type WodRounds struct {
	SetNumber          uint8  `json:"set_number" binding:"required"`
	RoundNumber        uint8  `json:"round_number" binding:"required"`
	RepetitionsByRound uint16 `json:"repetitions_by_round"`
	Time               uint32 `json:"time"`
	RestTime           uint32 `json:"rest_time"`
	RemainingTime      bool   `json:"remaining_Time"`
}

type ExerciseInWod struct {
	ExerciseId      uint16  `json:"exercise_id" binding:"required"`
	SetNumber       uint8   `json:"set_number" binding:"required"`
	RoundNumber     uint8   `json:"round_number" binding:"required"`
	Repetitions     uint16  `json:"repetitions" binding:"required"`
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
	WodTypeId uint8           `json:"wod_type_id" binding:"required"`
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

type CreateUserRequest struct {
	Id        string `json:"id" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
