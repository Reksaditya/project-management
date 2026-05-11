package models

import "time"

type BoardMember struct {
	BoardID  int64     `json:"board_id" db:"board_id" gorm:"column:board_id;primaryKey"`
	UserID   int64     `json:"user_id" db:"user_id" gorm:"column:user_id;primaryKey"`
	JoinedAt time.Time `json:"joined_at" db:"joined_at"`
}
