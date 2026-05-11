package models

import "github.com/google/uuid"

type Label struct {
	LabelID  int64     `json:"label_id" db:"label_id" gorm:"primaryKey;autoIncrement"`
	PublicID uuid.UUID `json:"public_id" db:"public_id"`
	Name     string    `json:"name" db:"name"`
	Color    string    `json:"color" db:"color"`
}
