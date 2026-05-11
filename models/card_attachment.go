package models

import (
	"time"

	"github.com/google/uuid"
)

type CardAttachment struct {
	CardAttID int64 `json:"card_att_id" db:"card_att_id" gorm:"primaryKey;autoIncrement"`
	PublicID  uuid.UUID `json:"public_id" db:"public_id"`
	File string `json:"file" db:"file"`
	UserID int64 `json:"user_id" db:"user_id"`
	CardID int64 `json:"card_id" db:"card_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}