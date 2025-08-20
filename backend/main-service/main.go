package main

import (
	"main-service/internal/database"
	"main-service/internal/link"
	"main-service/internal/logger"
	"main-service/internal/metrics"
	"net/http"

	_ "main-service/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title Swagger PseudoLink API
// @version 1.0
// @description Generated API-cellers on server

// @contact.url https://github.com/achi3v3
// @contact.email aamir-tutaev@mail.ru

// @host localhost:8080
// @BasePath /
func main() {

	redisClient := database.Init()
	defer redisClient.Close()

	linkRepo := link.NewLinkRepository(redisClient)
	linkService := link.NewLinkService(linkRepo)
	linkHandler := link.NewLinkHandler(linkService)

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	linkGroup := router.Group("/link")
	{
		linkGroup.POST("/create", linkHandler.Create)
		linkGroup.GET("/get", linkHandler.Get)
		linkGroup.DELETE("/delete", linkHandler.Delete)
	}
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Привет, Gin!")
	})
	logger.NewLogger()

	go func() {
		_ = metrics.Listen("127.0.0.1:8080")
	}()

	router.Run(":8080")
}
