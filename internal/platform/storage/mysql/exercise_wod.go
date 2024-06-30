package mysql

const sqlExerciseWodTable = "EXERCISE_WOD"

type sqlExerciseWod struct {
	ID                 string  `db:"id"`
	ExerciseID         string  `db:"exercise_id"`
	WodID              string  `db:"wod_id"`
	Reps               int32   `db:"repetitions"`
	Weight             float32 `db:"weight"`
	WeightUnit         string  `db:"weight_unit"`
	WodSection         int32   `db:"wod_section"`
	SectionTimerTypeID int32   `db:"section_timer_type_id"`
	SectionCap         int32   `db:"section_cap"`
}
