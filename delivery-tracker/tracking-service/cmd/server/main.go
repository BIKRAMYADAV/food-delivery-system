package main

import (
	"log"
	"tracking-service/internal/consumer"
	"tracking-service/internal/service"
)

func main() {
	consumer := consumer.NewConsumer()
	defer consumer.Close()
	trackingService := service.NewTrackingService()
	for{
		msg, err := consumer.Read()
		if err != nil{
			log.Println(err)
			continue
		}
		if err := trackingService.ProcessLocation(msg); err != nil{
			log.Println(err)
		}
	}
}	