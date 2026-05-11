package models

import (
	"github.com/Reksaditya/project-management/models/types"
	"github.com/google/uuid"
)

type CardPosition struct {
	CardPositionID int64           `json:"card_position_id" gorm:"primaryKey;autoIncrement"`
	PublicID       uuid.UUID       `json:"public_id" gorm:"type:uuid;not null"`
	ListID         int64           `json:"list_id" gorm:"column:list_id;not null"`
	CardOrder      types.UUIDArray `json:"card_order" gorm:"type:uuid[]"`
}
