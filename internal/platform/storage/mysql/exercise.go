package mysql

const (
	sqlExerciseTable = "EXERCISE"
)

type sqlExercise struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}
