package main
import "github.com/gin-gonic/gin"

import (
	"delivery-service/internal/service"
	"delivery-service/internal/handler"
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