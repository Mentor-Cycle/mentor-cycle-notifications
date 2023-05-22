package repositories

import (
	"context"
	"fmt"

	"github.com/Mentor-Cycle/mentor-cycle-notifications/api/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type NotificationRepository struct {
	conn *pgx.Conn
}

func NewNotificationRepository(conn *pgx.Conn) *NotificationRepository {
	return &NotificationRepository{conn: conn}
}

func (r *NotificationRepository) SaveNotification(message string) (models.Notification, error) {
	id := uuid.New()
	sql := "INSERT INTO notifications (id, message) VALUES ($1, $2)"

	_, err := r.conn.Exec(context.Background(), sql, id, message)
	if err != nil {
		return models.Notification{}, fmt.Errorf("error saving notification: %v", err)
	}

	return models.Notification{ID: id, Message: message}, nil
}

func (r *NotificationRepository) UpdateNotification(message string, id uuid.UUID) (models.Notification, error) {
	sql := "UPDATE notifications SET message = $1 WHERE id = $2"
	_, err := r.conn.Exec(context.Background(), sql, message, id)
	if err != nil {
		return models.Notification{}, fmt.Errorf("error saving notification: %v", err)
	}

	return models.Notification{ID: id, Message: message}, nil
}

func (r *NotificationRepository) GetNotifications() ([]models.Notification, error) {
	sql := "SELECT * FROM notifications"

	rows, err := r.conn.Query(context.Background(), sql)
	if err != nil {
		return nil, fmt.Errorf("error querying notifications: %v", err)
	}

	defer rows.Close()

	notifications := []models.Notification{}

	for rows.Next() {
		var id uuid.UUID
		var message string
		err = rows.Scan(&id, &message)
		if err != nil {
			return nil, fmt.Errorf("error scanning notification: %v", err)
		}
		notifications = append(notifications, models.Notification{ID: id, Message: message})
	}

	return notifications, nil
}
