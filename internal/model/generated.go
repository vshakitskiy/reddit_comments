// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CommentInput struct {
	PostID   string  `json:"postId"`
	Content  string  `json:"content"`
	ParentID *string `json:"parentId,omitempty"`
}

type Mutation struct {
}

type Query struct {
}

type Subscription struct {
}
