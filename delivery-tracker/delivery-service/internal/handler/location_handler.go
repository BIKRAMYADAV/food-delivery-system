package handler

import (
	"net/http"
	"tracking/internal/model"
	"tracking/internal/service"

	"github.com/gin-gonic/gin"
)

type LocationHandler struct {
	service *service.LocationService
}

func NewLocationHandler(service *service.LocationService) *LocationHandler{
	return &LocationHandler{
		service: service,
	}
}

func (h *LocationHandler) UpdateLocation(c *gin.Context){
	var location model.LocationUpdate

	if err := c.ShouldBindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.ProcessLocation(location); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : "unable to process location",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "location recieved", 
	})
}