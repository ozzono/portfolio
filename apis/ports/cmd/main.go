package main

import (
	"context"
	"ports/internal/server"
	"ports/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	Version            = "1.0.0"
	PostgresqlUser     = "postgres"
	PostgresqlPassword = "postgres"
	PostgresqlHost     = "localhost"
	PostgresqlPort     = "5432"
	UserPostgresqlDB   = "ports_service"
	PostgresqlSSLMode  = "disable"
)

func main() {
	logger := log.New().With(context.TODO(), "version", Version)
	dsn := "host=localhost user=postgres password=postgres dbname=ports_service port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("gorm.Open err", err)
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			logger.Error(err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			logger.Error(err)
		}
	}()

	s := server.NewServer(logger, db)
	logger.Info("starting server")
	logger.Fatal(s.Run())
}
