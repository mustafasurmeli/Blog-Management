package models

type Follow struct {
	ID         uint `json:"id" gorm:"primaryKey"`
	FollowerID uint `json:"follower_id"`
	FollowedId uint `json:"followed_id"`
	Follower   User `json:"follower" gorm:"foreignKey:FollowerID"`
	Followed   User `json:"followed" gorm:"foreignKey:FollowedID"`
}
