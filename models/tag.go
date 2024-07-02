package models

type Tag struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"unique;not null"`
	Posts []Post `json:"posts" gorm:"many2many:post_tags"`
}
