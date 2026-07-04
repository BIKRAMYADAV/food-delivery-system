package main
import "github.com/gin-gonic/gin"

import (
	"tracking/internal/service"
	"tracking/internal/handler"
)

func main(){
	router := gin.Default()
	locationService := service.NewLocationService()
	locationHandler := handler.NewLocationHandler(locationService)
	router.GET("/health", func (c *gin.Context){
		c.JSON(200, gin.H{
			"status":"running",
		})
	})
	router.POST("/location", locationHandler.UpdateLocation)
	router.Run(":8080")
}