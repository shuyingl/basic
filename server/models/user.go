package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID `json:"id"           gorm:"type:UUID;primaryKey;"`
	Email string    `json:"email"        gorm:"type:TEXT;NOT NULL;UNIQUE"`
}
