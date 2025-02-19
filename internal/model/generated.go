// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AuthPayload struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}

type CommentInput struct {
	PostID   int32  `json:"postId"`
	Content  string `json:"content"`
	ParentID *int32 `json:"parentId,omitempty"`
}

type CommentsConnection struct {
	Meta *ConnectionMeta `json:"meta,omitempty"`
	Rows []*Comment      `json:"rows,omitempty"`
}

type ConnectionMeta struct {
	TotalRow   int32 `json:"totalRow"`
	TotalPages int32 `json:"totalPages"`
	Limit      int32 `json:"limit"`
	Page       int32 `json:"page"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Mutation struct {
}

type PaginationInput struct {
	Limit int32 `json:"limit"`
	Page  int32 `json:"page"`
}

type PostInput struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	CommentsEnabled bool   `json:"commentsEnabled"`
}

type PostsConnection struct {
	Meta *ConnectionMeta `json:"meta,omitempty"`
	Rows []*Post         `json:"rows,omitempty"`
}

type Query struct {
}

type RegisterInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Subscription struct {
}
