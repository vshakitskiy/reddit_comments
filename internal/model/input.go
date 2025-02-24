package model

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PostInput struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	CommentsEnabled bool   `json:"commentsEnabled"`
}

type RegisterInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
