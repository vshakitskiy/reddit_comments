package model

import "time"

type Post struct {
	ID              string `gorm:"primaryKey" json:"id"`
	Title           string `gorm:"not null" json:"title"`
	Description     string `gorm:"not null" json:"description"`
	CommentsEnabled bool   `json:"commentsEnabled"`
	TotalComments   int32  `json:"totalComments"`
	UserID          string `gorm:"not null;index" json:"userId"`

	User     *User               `gorm:"foreignKey:UserID" json:"user"`
	Comments *CommentsConnection `gorm:"-" json:"comments"`

	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null" json:"updatedAt"`
}
