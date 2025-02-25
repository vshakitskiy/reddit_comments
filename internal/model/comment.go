package model

import "time"

type Comment struct {
	ID       string  `gorm:"primaryKey" json:"id"`
	Content  string  `gorm:"not null" json:"content"`
	UserID   string  `gorm:"not null;index" json:"userId"`
	ParentID *string `gorm:"index" json:"parentId,omitempty"`
	PostID   string  `gorm:"not null;index" json:"postId"`

	User    *User               `gorm:"-" json:"user"`
	Parent  *Comment            `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	Replies *CommentsConnection `gorm:"-" json:"replies"`

	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null" json:"updatedAt"`
}

func (u *Comment) GetID() string {
	return u.ID
}
