package mysql

const (
	sqlExerciseTable = "exercise"
)

type sqlExercise struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}
