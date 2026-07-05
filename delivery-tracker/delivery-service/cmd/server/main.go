package main
import "github.com/gin-gonic/gin"

import (
	"delivery-service/pkg/kafka"
	"delivery-service/internal/config"
	"delivery-service/internal/service"
	"delivery-service/internal/handler"
)

func main(){
	cfg := config.Load()

	producer := kafka.NewProducer(
		cfg.KafkaBroker,
		cfg.KafkaTopic,
	)
	defer producer.Close()

	router := gin.Default()
	locationService := service.NewLocationService(producer)
	locationHandler := handler.NewLocationHandler(locationService)
	router.GET("/health", func (c *gin.Context){
		c.JSON(200, gin.H{
			"status":"running",
		})
	})
	router.POST("/location", locationHandler.UpdateLocation)
	router.Run(":"+cfg.ServerPort)
}