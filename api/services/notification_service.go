package services

import (
	"fmt"

	"github.com/Mentor-Cycle/mentor-cycle-notifications/api/dtos"
	"github.com/Mentor-Cycle/mentor-cycle-notifications/api/repositories"
	"github.com/google/uuid"
)

type NotificationService struct {
	repository *repositories.NotificationRepository
}

func NewNotificationService(repository *repositories.NotificationRepository) *NotificationService {
	return &NotificationService{repository: repository}
}

func (s *NotificationService) SaveNotification(message string) (dtos.Response, error) {
	data, err := s.repository.SaveNotification(message)
	if err != nil {
		return dtos.Response{Success: false, Message: err.Error()}, err
	}
	fmt.Println("Service =>" + data.Message)
	return dtos.Response{Success: true, Data: data, Message: data.Message}, nil

}

func (s *NotificationService) GetNotifications() (dtos.Response, error) {
	data, err := s.repository.GetNotifications()
	if err != nil {
		return dtos.Response{Success: false, Message: err.Error()}, err
	}

	return dtos.Response{Success: true, Data: data}, nil
}

func (s *NotificationService) UpdateNotification(message string, id uuid.UUID) (dtos.Response, error) {
	data, err := s.repository.UpdateNotification(message, id)
	if err != nil {
		return dtos.Response{Success: false, Message: err.Error()}, err
	}

	return dtos.Response{Success: true, Data: data}, nil
}
