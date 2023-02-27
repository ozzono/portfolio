package main

import (
	"rover/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "rover/docs"
)

// @title Mars rover control panel
// @version 2.0
// @description This is a Mars rover control panel

// @host localhost:8080
// @BasePath /v2
func main() {
	r := gin.Default()

	v2 := r.Group("/v2")
	v2.GET("/plateau/set", handler.SetPlateau)
	v2.GET("/plateau/show", handler.ShowPlateau)
	v2.PUT("/rover", handler.LandRover)
	v2.GET("/rover/:id", handler.MoveRover)

	url := ginSwagger.URL("doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
