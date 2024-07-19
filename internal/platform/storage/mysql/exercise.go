package mysql

const (
	sqlExerciseTable = "EXERCISE"
)

type sqlExercise struct {
	Name string `db:"name"`
}
