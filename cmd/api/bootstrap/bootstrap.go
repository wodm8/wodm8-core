package bootstrap

import (
	"context"
	"github.com/wodm8/wodm8-core/initializers"
	"github.com/wodm8/wodm8-core/internal/application"
	"github.com/wodm8/wodm8-core/internal/platform/server"
	storage "github.com/wodm8/wodm8-core/internal/platform/storage/mysql"
	_ "gorm.io/driver/mysql"
)

func Run() error {

	initializers.LoadEnvVariables()
	initializers.ConnectToDb()

	cfg := initializers.Cfg

	db := initializers.DB

	exerciseRepository := storage.NewExerciseRepository(db)
	wodRepository := storage.NewWodRepository(db)
	exerciseWodRepository := storage.NewExerciseWodRepository(db)
	wodSetRepository := storage.NewWodSetRepository(db)
	wodRoundRepository := storage.NewWodRoundRepository(db)
	usersRepository := storage.NewUsersRepository(db)
	membersRepository := storage.NewMembersRepository(db)

	exerciseService := application.NewExerciseService(exerciseRepository)

	wodService := application.NewWodService(wodRepository, wodSetRepository, wodRoundRepository, exerciseWodRepository)

	usersService := application.NewUsersService(usersRepository, membersRepository)

	membersService := application.NewMemberService(membersRepository)

	ctx, srv := server.New(
		context.Background(),
		cfg.HostServer,
		cfg.PortServer,
		cfg.ShutdownTimeout,
		wodService,
		exerciseService,
		usersService,
		membersService,
	)
	return srv.Run(ctx)
}
