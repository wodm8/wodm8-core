package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wodm8/wodm8-core/internal/platform/handler/health"
	"log"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine
}

func NewServer(host string, port uint) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),
	}
	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Printf("Starting server at %s", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
}
