package model

import "time"

type Post struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	CommentsEnabled bool   `json:"commentsEnabled"`
	TotalComments   int    `json:"totalComments"`
	UserID          string `json:"userId"`

	User     *User               `json:"user"`
	Comments *CommentsConnection `json:"comments"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
