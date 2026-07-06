package service

import (
	"encoding/json"
	"log"
	"tracking-service/internal/model"
	"github.com/segmentio/kafka-go"
)
type TrackingService struct{
	
}
func NewTrackingService() *TrackingService{
	return &TrackingService{}
}

func (s* TrackingService) ProcessLocation(msg kafka.Message) error{
	var location model.LocationUpdate
	if err := json.Unmarshal(msg.Value, &location); err != nil {
		return err 
	}
	log.Printf(
		"Order %s Rider (%f,%f)",
		location.OrderID,
		location.Latitude,
		location.Longitude,
	)
	return nil
}