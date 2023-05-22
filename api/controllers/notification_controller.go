package controllers

import (
	"net/http"

	"github.com/Mentor-Cycle/mentor-cycle-notifications/api/services"
	"github.com/Mentor-Cycle/mentor-cycle-notifications/api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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
	data, err := c.service.SaveNotification(request.Message)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (c *NotificationController) GetNotifications(ctx *gin.Context) {
	data, err := c.service.GetNotifications()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (c *NotificationController) UpdateNotification(ctx *gin.Context) {
	idSrt := ctx.Param("id")
	id, err := uuid.Parse(idSrt)
	if err != nil {
		helpers.GenerateCtxError(ctx, err.Error(), http.StatusBadRequest)
	}
	var request SaveNotificationRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		helpers.GenerateCtxError(ctx, err.Error(), http.StatusBadRequest)
	}
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		helpers.GenerateCtxError(ctx, err.Error(), http.StatusBadRequest)
	}

	data, err := c.service.UpdateNotification(request.Message, id)

	if err != nil {
		helpers.GenerateCtxError(ctx, err.Error(), http.StatusBadRequest)
		return
	}
	ctx.JSON(http.StatusOK, data)
}
