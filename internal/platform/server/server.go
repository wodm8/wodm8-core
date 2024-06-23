package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	crossfit "github.com/wodm8/wodm8-core/internal"
	"github.com/wodm8/wodm8-core/internal/platform/server/handler/exercise"
	"github.com/wodm8/wodm8-core/internal/platform/server/handler/health"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	//deps
	exerciseRepository crossfit.ExerciseRepository
}

func NewServer(host string, port int, exerciseRepository crossfit.ExerciseRepository) Server {
	srv := Server{
		engine:             gin.New(),
		httpAddr:           fmt.Sprintf("%s:%d", host, port),
		exerciseRepository: exerciseRepository,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Printf("Start server at %s", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/api/v1/exercises", exercise.CreateHandler(s.exerciseRepository))
}
