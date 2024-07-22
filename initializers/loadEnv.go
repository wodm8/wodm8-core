package initializers

import (
	"github.com/kelseyhightower/envconfig"
	"log"
	"time"
)

var Cfg Config

func LoadEnvVariables() {
	err := envconfig.Process("WODM8", &Cfg)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

type Config struct {
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
	JwtSecret  []byte
}
