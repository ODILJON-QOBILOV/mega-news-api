package main

import (
	"github.com/gin-gonic/gin"
	"github.com/newsapi/v2/config"
	"github.com/newsapi/v2/routes"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/files" 
	_ "github.com/newsapi/v2/docs"
	"github.com/gin-contrib/cors"
)
// @title News API
// @version 1.0
// @description News API
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	config.ConnectDB()
	router := gin.Default()
	router.Use(cors.Default())
	
	router.GET("/", RootAPI)


    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.RegisterRoutes(router)

	router.Run(":8080")
}

func RootAPI(c *gin.Context) {
	c.JSON(200, gin.H{"message": "root route"})
}