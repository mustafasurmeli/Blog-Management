package models

type Like struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	PostID uint `json:"post_id"`
	UserID uint `json:"user_id"`
	Post   Post `json:"post" gorm:"foreignKey:PostID"`
	User   User `json:"user" gorm:"foreignKey:UserID"`
}
