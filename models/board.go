package models

import (
	"time"

	"github.com/google/uuid"
)

type Board struct {
	BoardID       int64      `json:"board_id" db:"board_id" gorm:"primaryKey;autoIncrement"`
	PublicID      uuid.UUID  `json:"public_id" db:"public_id"`
	Title         string     `json:"title" db:"title"`
	Description   string     `json:"description" db:"description"`
	OwnerID       int64      `json:"owner_id" db:"owner_id" gorm:"column:owner_id"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	Duedate       *time.Time `json:"due_date,omitempty" db:"due_date"`
	OwnerPublicID uuid.UUID  `json:"owner_public_id" db:"owner_public_id"`
}
