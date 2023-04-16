package repositories

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type NotificationRepository struct {
    conn *pgx.Conn
}

func NewNotificationRepository(conn *pgx.Conn) *NotificationRepository {
    return &NotificationRepository{conn: conn}
}

func (r *NotificationRepository) SaveNotification(message string) error {
    id := uuid.New()
    sql := "INSERT INTO notifications (id, message) VALUES ($1, $2)"

    _, err := r.conn.Exec(context.Background(), sql, id, message)
    if err != nil {
        return fmt.Errorf("error saving notification: %v", err)
    }

    return nil
}
