package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"ports/internal/repository"
	rest "ports/internal/rest"
	"ports/log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	ServerPort = ":8000"
)

// server
type server struct {
	log log.Logger
	db  *gorm.DB
	gin *gin.Engine
}

func NewServer(
	log log.Logger,
	db *gorm.DB,
) *server {
	return &server{
		log: log,
		db:  db,
		gin: gin.New(),
	}
}

func (s server) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	group := s.gin.Group("/challenge")
	repo := repository.NewRepository(s.db, s.log)
	svc := repository.NewService(repo, s.log)

	handlers := rest.NewPortHandlers(group, s.log, svc)
	handlers.MapPortRoutes()

	go func() {
		s.runHttpServer()
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		s.log.Errorf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		s.log.Errorf("ctx.Done: %v", done)
	}

	s.log.Info("Server Exited Properly")
	return nil
}

func (s *server) runHttpServer() {
	s.gin.GET("/ping", func(c *gin.Context) {
		s.log.Info("pong")
		c.String(http.StatusOK, "pong")
	})

	go func() {
		if err := s.gin.Run(ServerPort); err != nil {
			s.log.Fatalf("Error starting TLS Server: ", err)
		}
	}()
}
