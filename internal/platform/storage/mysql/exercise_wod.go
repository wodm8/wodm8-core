package mysql

const sqlExerciseWodTable = "EXERCISE_WOD"

type sqlExerciseWod struct {
	ID              string  `db:"id"`
	WodID           string  `db:"wod_id"`
	ExerciseID      int32   `db:"exercise_id"`
	RoundNumber     int32   `db:"round_number"`
	SetNumber       int32   `db:"set_number"`
	Repetitions     int32   `db:"repetitions"`
	RepetitionsUnit string  `db:"repetitions_unit"`
	Weight          float32 `db:"weight"`
	WeightUnit      string  `db:"weight_unit"`
	Distance        float32 `db:"distance"`
	DistanceUnit    string  `db:"distance_unit"`
}
