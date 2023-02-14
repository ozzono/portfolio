package main

import (
	"flag"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	_ "car-rental/docs"
	"car-rental/internal/handler"
	"car-rental/internal/repository"
	"car-rental/route"
)

// @title Car rental API
// @description golang api - Car rental API

// @host localhost:8080
// @BasePath /api/v1
func main() {
	debug := flag.Bool("debug", false, "enables debug mode")
	log := zap.NewExample().Sugar()
	dbClient, err := repository.NewDBClient(*debug)
	if err != nil {
		log.Errorf("repository.NewDBClient - %v", err)
		return
	}
	handler := handler.NewHandler(dbClient, log, *debug)
	r := route.Routes(handler, *debug)
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run(":8080")
}
