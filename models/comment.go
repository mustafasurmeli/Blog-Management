package models

import "time"

type Comment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Content   string    `json:"content" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	PostID    uint      `json:"post_id"`
	UserID    uint      `json:"user_id"`
	Post      Post      `json:"post" gorm:"foreignKey:PostID"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
}
