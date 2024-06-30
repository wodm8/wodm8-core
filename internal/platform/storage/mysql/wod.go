package mysql

const (
	sqlWodTable = "WOD"
)

type sqlWod struct {
	ID             string `db:"id"`
	Name           string `db:"wod_name"`
	Rounds         int32  `db:"rounds"`
	NumberSections int32  `db:"number_sections"`
	TimerTypeId    int32  `db:"timer_type_id"`
}
