package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/wodm8/wodm8-core/internal/domain"
	"log"
	"reflect"
	"time"

	"github.com/huandu/go-sqlbuilder"
	"github.com/wodm8/wodm8-core/internal/crossfit"
)

type WodRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

func NewWodRepository(db *sql.DB, dbTimeout time.Duration) *WodRepository {
	return &WodRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

func (r *WodRepository) Save(ctx context.Context, wod crossfit.Wod) error {
	wodSQLStruct := sqlbuilder.NewStruct(new(sqlWod))
	query, args := wodSQLStruct.InsertInto(sqlWodTable, sqlWod{
		ID:        wod.ID().String(),
		Name:      wod.Name().String(),
		WodDate:   wod.Date().String(),
		WodTypeId: wod.TypeID().Int(),
	}).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	fmt.Printf("query wod: %+v\n", query)
	fmt.Printf("ctx timeout: %+v\n", ctxTimeout)
	fmt.Printf("args:%+v\n", args)
	_, err := r.db.ExecContext(ctxTimeout, query, args...)

	if err != nil {
		fmt.Printf("error saving wod: %v", err)
		return err
	}
	return nil
}

func (r *WodRepository) Get(ctx context.Context, id string) ([]domain.WodResult, error) {
	var wodSlice = make([]domain.WodResult, 0)
	sb := sqlbuilder.NewSelectBuilder()
	q, args := sb.Select("*").
		From("WOD_CREATED").
		Where(sb.EQ("wod_id", id)).Build()

	res, err := r.db.QueryContext(ctx, q, args...)
	if err != nil {
		fmt.Printf("error getting wod_data: %v", err)
	}
	defer res.Close()

	for res.Next() {
		var wodResult domain.WodResult
		s := reflect.ValueOf(&wodResult).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}
		var err = res.Scan(columns...)
		if err != nil {
			log.Fatal(err)
		}
		wodSlice = append(wodSlice, wodResult)
	}

	return wodSlice, err
}
