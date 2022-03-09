package main

import (
	"context"
	"ports/internal/server"
	"ports/pkg/log"
	"ports/pkg/postgres"
)

const (
	Version    = "1.0.0"
	pgUser     = "postgres"
	pgPassword = "postgres"
	pgHost     = "ports_postgresql"
	pgPort     = "5432"
	pgDBName   = "ports_service"
	pgSSLMode  = "disable"
)

func main() {
	logger := log.New().With(context.TODO(), "version", Version)

	pgxConn, err := postgres.NewPgxConn(postgres.Config{
		Host:     pgHost,
		Port:     pgPort,
		User:     pgUser,
		DBNAme:   pgDBName,
		SSLMode:  pgSSLMode,
		Password: pgPassword,
	})
	if err != nil {
		logger.Fatal("cannot connect to postgres", err)
	}
	defer pgxConn.Close()

	s := server.NewServer(logger, pgxConn)
	logger.Info("starting server")
	logger.Fatal(s.Run())
}
