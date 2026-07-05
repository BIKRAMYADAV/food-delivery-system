package service

import (
	"delivery-service/internal/model"
	"delivery-service/pkg/kafka"
	"encoding/json"
	"log"
)

type LocationService struct {
	producer *kafka.Producer
}

func NewLocationService(producer *kafka.Producer) *LocationService {
	return &LocationService{
		producer: producer,
	}
}

func (s *LocationService) ProcessLocation(location model.LocationUpdate) error {
	data, err := json.Marshal(location)
	if(err != nil){
		log.Printf("Kafka publish failed: %v", err)
		return err
	}
	log.Printf("delivery_partner=%s lat=%f lon=%f",
		location.DeliveryPartnerID,
		location.Latitude,
		location.Longitude,
	)
	log.Println("Published successfully")
	// return nil
	return s.producer.Publish(data)
}
