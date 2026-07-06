package model

import "time"

type LocationUpdate struct {
	DeliveryPartnerID string `json:"delivery_partner_id" binding:"required"`
	OrderID		   string `json:"order_id" binding:"required"`
	Latitude	   float64 `json:"latitude" binding:"required"`
	Longitude 	   float64 `json:"longitude" binding:"required"`
	Speed		   float64 `json:"speed" `
	Heading		   float64 `json:"heading"`
	Timestamp	   time.Time `json:"timestamp"`
}