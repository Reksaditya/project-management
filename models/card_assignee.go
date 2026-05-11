package models

type CardAssignee struct {
	CardID int64 `json:"card_id" db:"card_id" gorm:"column:card_id"`
	UserID int64 `json:"user_id" db:"user_id" gorm:"column:user_id"`
}