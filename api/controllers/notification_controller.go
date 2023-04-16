package controllers

import (
	"net/http"

	"github.com/Mentor-Cycle/mentor-cycle-notifications/api/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SaveNotificationRequest struct {
	Message string `json:"message" binding:"required"`
}

type NotificationController struct {
	service *services.NotificationService
}

func NewNotificationController(service *services.NotificationService) *NotificationController {
	return &NotificationController{service: service}
}

func (c *NotificationController) SaveNotification(ctx *gin.Context) {
	var request SaveNotificationRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.SaveNotification(request.Message)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "notification saved"})
}
