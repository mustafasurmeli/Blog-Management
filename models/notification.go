package models

import "time"

type Notification struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Content   string    `json:"content" gorm:"not null"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
}
