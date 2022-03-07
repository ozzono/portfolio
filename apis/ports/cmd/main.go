package main

import (
	"context"
	"ports/internal/server"
	"ports/log"
	"ports/pkg/postgres"
)

const (
	Version            = "1.0.0"
	PostgresqlUser     = "postgres"
	PostgresqlPassword = "postgres"
	PostgresqlHost     = "localhost"
	PostgresqlPort     = "5432"
	PostgresqlDBName   = "ports_service"
	PostgresqlSSLMode  = "disable"
)

func main() {
	logger := log.New().With(context.TODO(), "version", Version)

	pgxConn, err := postgres.NewPgxConn(postgres.Config{
		Host:     PostgresqlUser,
		Port:     PostgresqlPassword,
		User:     PostgresqlHost,
		DBNAme:   PostgresqlPort,
		SSLMode:  PostgresqlDBName,
		Password: PostgresqlSSLMode,
	})
	if err != nil {
		logger.Fatal("cannot connect to postgres", err)
	}
	defer pgxConn.Close()

	s := server.NewServer(logger, pgxConn)
	logger.Info("starting server")
	logger.Fatal(s.Run())
}
