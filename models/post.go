package models

import (
	"time"
)

type Post struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	Title      string     `json:"title" gorm:"not null"`
	Content    string     `json:"content" gorm:"not null"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	UserID     uint       `json:"user_id"`
	User       User       `json:"user" gorm:"foreignKey:UserID"`
	Comments   []Comment  `json:"comments"`
	Likes      []Like     `json:"likes"`
	Categories []Category `json:"categories" gorm:"many2many:post_categories"`
	Tags       []Tag      `json:"tags" gorm:"many2many:post_tags"`
}
