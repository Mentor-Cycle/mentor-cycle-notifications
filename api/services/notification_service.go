package services

import (
	"github.com/Mentor-Cycle/mentor-cycle-notifications/api/repositories"
)

type NotificationService struct {
    repository *repositories.NotificationRepository
}

func NewNotificationService(repository *repositories.NotificationRepository) *NotificationService {
    return &NotificationService{repository: repository}
}

func (s *NotificationService) SaveNotification(message string) error {
    return s.repository.SaveNotification(message)
}
