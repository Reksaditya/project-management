package models

import (
	"time"

	"github.com/google/uuid"
)

type Card struct {
	CardID      int64      `json:"card_id" db:"card_id" gorm:"primaryKey;autoIncrement"`
	PublicID    uuid.UUID  `json:"public_id" db:"public_id"`
	ListID      int64      `json:"list_id" db:"list_id" gorm:"column:list_id"`
	Title       string     `json:"title" db:"title"`
	Description string     `json:"description" db:"description"`
	Duedate     *time.Time `json:"due_date,omitempty" db:"due_date"`
	Position    int        `json:"position" db:"position"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
}
