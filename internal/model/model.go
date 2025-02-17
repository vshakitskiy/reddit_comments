package model

import "time"

type UserMemory struct {
	ID           string
	Username     string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type PostMemory struct {
	ID              string
	Title           string
	Description     string
	CommentsEnabled bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
	UserID          string
	CommentCount    int

	User *UserMemory
}

type CommentMemory struct {
	ID        string
	Content   string
	ParentID  *string
	PostID    string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time

	User    *UserMemory
	Parent  *CommentMemory
	Replies []*CommentMemory
}
