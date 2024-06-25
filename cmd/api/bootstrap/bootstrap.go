package bootstrap

import "github.com/wodm8/wodm8-core/internal/platform/server"

const (
	host = "localhost"
	port = 8080
)

func Run() error {
	srv := server.NewServer(host, port)
	return srv.Run()
}
