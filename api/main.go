package main

import (
	"log"

	"github.com/Mentor-Cycle/mentor-cycle-notifications/api/controllers"
	"github.com/Mentor-Cycle/mentor-cycle-notifications/api/database"
	"github.com/Mentor-Cycle/mentor-cycle-notifications/api/repositories"
	"github.com/Mentor-Cycle/mentor-cycle-notifications/api/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main () {
	router := gin.Default()
	conn, err := database.Connect()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	repository := repositories.NewNotificationRepository(conn)
	service := services.NewNotificationService(repository)
	setupRoutes(router, service)
	router.Run(":8080")
}

func setupRoutes(router *gin.Engine, service *services.NotificationService) {

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // You can replace * with your client URL
	router.Use(cors.New(config))

	controller := controllers.NewNotificationController(service)

	router.POST("/notifications", controller.SaveNotification)
}

