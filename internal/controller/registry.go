package controller

import (
	"Building/docs"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func Registry() {
	r := gin.Default()

	docs.SwaggerInfo.Title = "BuildingDB API"
	docs.SwaggerInfo.Description = "add/get buildings"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"

	r.POST("/CreateBuilding", CreateBuilding)
	r.GET("/GetBuildings", GetBuildings)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	port := viper.GetString("port")
	err := r.Run(port)
	if err != nil {
		log.Fatalf("Failed to start the web server - Error: %v", err)
	}
}
