package main

import (
	"main-service/backend/main-service/internal/database"
	"main-service/backend/main-service/internal/link"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	redisClient := database.Init()
	defer redisClient.Close()

	linkRepo := link.NewLinkRepository(redisClient)
	linkService := link.NewLinkService(linkRepo)
	linkHandler := link.NewLinkHandler(linkService)

	router := gin.Default()

	linkGroup := router.Group("/link")
	{
		linkGroup.POST("/create", linkHandler.Create)
		linkGroup.GET("/get", linkHandler.Get)
		linkGroup.DELETE("/delete", linkHandler.Delete)
	}
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Привет, Gin!")
	})
	router.Run(":8080")
}
