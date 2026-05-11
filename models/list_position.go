package models

import (
	"github.com/Reksaditya/project-management/models/types"
	"github.com/google/uuid"
)

type ListPosition struct {
	ListID    int64           `json:"list_id" db:"list_id" gorm:"primaryKey;autoIncrement"`
	PublicID  uuid.UUID       `json:"public_id" db:"public_id" gorm:"public_id"`
	BoardID   int64           `json:"board_id" db:"board_id" gorm:"column:board_id"`
	ListOrder types.UUIDArray `json:"list_order"`
}
