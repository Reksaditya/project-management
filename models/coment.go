package models

import (
	"time"

	"github.com/google/uuid"
)

type Coment struct {
	ComentID     int64     `json:"coment_id" db:"coment_id" gorm:"primaryKey;autoIncrement"`
	PublicID     uuid.UUID `json:"public_id" db:"public_id"`
	CardID       int64     `json:"card_id" db:"card_id"`
	CardPublicID uuid.UUID `json:"card_public_id" db:"card_public_id"`
	UserID       int64     `json:"user_id" db:"user_id"`
	UserPublicID uuid.UUID `json:"user_public_id" db:"user_public_id"`
	Message      string    `json:"message" db:"message"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}
