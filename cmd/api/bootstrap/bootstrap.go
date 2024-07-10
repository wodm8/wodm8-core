package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/wodm8/wodm8-core/internal/creating"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
	"github.com/wodm8/wodm8-core/internal/platform/server"
	"github.com/wodm8/wodm8-core/internal/platform/storage/mysql"
)

func Run() error {
	var cfg config
	err := envconfig.Process("WODM8", &cfg)
	if err != nil {
		return err
	}

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		log.Fatal("Error db connection", err)
	}

	exerciseRepository := mysql.NewExerciseRepository(db, cfg.DbTimeout)
	wodRepository := mysql.NewWodRepository(db, cfg.DbTimeout)
	exerciseWodRepository := mysql.NewExerciseWodRepository(db, cfg.DbTimeout)

	exerciseService := creating.NewExerciseService(exerciseRepository)
	wodService := creating.NewWodService(wodRepository, exerciseWodRepository)

	ctx, srv := server.New(context.Background(), cfg.HostServer, cfg.PortServer, cfg.ShutdownTimeout, wodService, exerciseService)
	return srv.Run(ctx)
}

type config struct {
	//Server config
	HostServer      string
	PortServer      int
	ShutdownTimeout time.Duration `default:"10s"`
	//DB config
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     int
	DbName     string
	DbTimeout  time.Duration `default:"5s"`
}
