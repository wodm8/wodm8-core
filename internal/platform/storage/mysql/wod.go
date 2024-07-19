package mysql

const (
	sqlWodTable = "WOD"
)

type sqlWod struct {
	ID        string `db:"id"`
	Name      string `db:"name"`
	WodDate   string `db:"wod_date"`
	WodTypeId int32  `db:"wod_type_id"`
}
