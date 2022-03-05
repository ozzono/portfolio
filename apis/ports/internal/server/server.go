package server

import (
	"net/http"

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

func (s server) Run() {

	group := s.gin.Group("/challenge")
	repo := repository.NewRepository(s.db, s.log)
	svc := repository.NewService(repo, s.log)

	handlers := rest.NewPortHandlers(group, s.log, svc)
	handlers.MapPortRoutes()

	go func() {
		s.runHttpServer()
	}()
}

func (s *server) runHttpServer() {
	s.gin.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	go func() {
		if err := s.gin.Run(ServerPort); err != nil {
			s.log.Fatalf("Error starting TLS Server: ", err)
		}
	}()
}
