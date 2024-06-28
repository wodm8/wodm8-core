package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/wodm8/wodm8-core/commons"
	"github.com/wodm8/wodm8-core/internal/platform/server"
	"github.com/wodm8/wodm8-core/internal/platform/storage/mysql"
)

func Run() error {
	err := godotenv.Load()

	var (
		hostServer      = os.Getenv("HOST_SERVER")
		portServer      = os.Getenv("PORT_SERVER")
		dbUser          = os.Getenv("DB_USER")
		dbPswd          = os.Getenv("DB_PASSWORD")
		dbHost          = os.Getenv("DB_HOST")
		dbPort          = os.Getenv("DB_PORT")
		dbName          = os.Getenv("DB_NAME")
		shutdownTimeout = 10 * time.Second
		dbTimeout       = 5 * time.Second
	)

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPswd, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		log.Fatal("Error db connection", err)
	}

	exerciseRepository := mysql.NewExerciseRepository(db, dbTimeout)

	portSrv, err := commons.GetenvInt(portServer)
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	ctx, srv := server.New(context.Background(), hostServer, portSrv, shutdownTimeout, exerciseRepository)
	return srv.Run(ctx)
}
