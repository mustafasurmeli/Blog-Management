package models

type Category struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"unique;not null"`
	Description string `json:"description"`
	Posts       []Post `json:"posts" gorm:"many2many:post_categories"`
}
