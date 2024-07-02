package models

type Profile struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	UserID         uint   `json:"userID"`
	Bio            string `json:"bio"`
	ProfilePicture string `json:"profilePicture"`
}
