package service

import (
	"log"
	"tracking/internal/model"
)

type LocationService struct{}

func NewLocationService() *LocationService {
	return &LocationService{}
}

func (s *LocationService) ProcessLocation(location model.LocationUpdate) error {
	log.Printf("delivery_partner=%s lat=%f lon=%f",
		location.DeliveryPartnerID,
		location.Latitude,
		location.Longitude,
	)
	return nil
}
