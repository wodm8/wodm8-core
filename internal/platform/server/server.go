package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wodm8/wodm8-core/internal/application"
	"github.com/wodm8/wodm8-core/internal/platform/server/handler/exercise"
	"github.com/wodm8/wodm8-core/internal/platform/server/handler/health"
	"github.com/wodm8/wodm8-core/internal/platform/server/handler/wod"
)

type Server struct {
	httpAddr        string
	engine          *gin.Engine
	shutdownTimeout time.Duration

	wodService      application.WodService
	exerciseService application.ExerciseService
}

func New(ctx context.Context, host string, port int, shutdownTimeout time.Duration, wodService application.WodService, exerciseService application.ExerciseService) (context.Context, Server) {
	srv := Server{
		engine:          gin.Default(),
		httpAddr:        fmt.Sprintf("%s:%d", host, port),
		shutdownTimeout: shutdownTimeout,

		wodService:      wodService,
		exerciseService: exerciseService,
	}

	srv.registerRoutes()
	return serverContext(ctx), srv
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/api/v1/exercises", exercise.CreateHandler(s.exerciseService))
	s.engine.POST("/api/v1/wod", wod.CreateWodHandler(s.wodService))
	s.engine.GET("/api/v1/wod", wod.GetWodHandler(s.wodService))
}

func (s *Server) Run(ctx context.Context) error {
	log.Printf("Start server at %s", s.httpAddr)

	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shut down", err)
		}
	}()

	<-ctx.Done()
	ctxShutDown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()
	return srv.Shutdown(ctxShutDown)
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()

	return ctx
}
