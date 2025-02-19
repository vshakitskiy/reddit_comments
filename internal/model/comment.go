package model

import "time"

type Comment struct {
	ID       string  `json:"id"`
	Content  string  `json:"content"`
	UserId   string  `json:"userId"`
	ParentId *string `json:"parentId"`

	User    *User               `json:"user"`
	Parent  *Comment            `json:"parent,omitempty"`
	Replies *CommentsConnection `json:"replies"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
