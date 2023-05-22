package models

import "github.com/google/uuid"

type Notification struct {
	ID	    uuid.UUID `json:"id"`
	Message string `json:"message"`
}