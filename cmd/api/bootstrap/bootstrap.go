package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/wodm8/wodm8-core/internal/application"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/kelseyhightower/envconfig"
	"github.com/wodm8/wodm8-core/internal/platform/server"
	storage "github.com/wodm8/wodm8-core/internal/platform/storage/mysql"
)

func Run() error {
	var cfg config
	err := envconfig.Process("WODM8", &cfg)
	if err != nil {
		return err
	}

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)
	db, err := gorm.Open(mysql.Open(mysqlURI), &gorm.Config{})
	if err != nil {
		log.Fatal("Error db connection", err)
	}

	exerciseRepository := storage.NewExerciseRepository(db)
	wodRepository := storage.NewWodRepository(db)
	exerciseWodRepository := storage.NewExerciseWodRepository(db)
	wodSetRepository := storage.NewWodSetRepository(db)
	wodRoundRepository := storage.NewWodRoundRepository(db)

	exerciseService := application.NewExerciseService(exerciseRepository)

	wodService := application.NewWodService(wodRepository, wodSetRepository, wodRoundRepository, exerciseWodRepository)

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
