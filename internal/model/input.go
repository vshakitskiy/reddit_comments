package model

import "github.com/vshakitskiy/reddit_comments/pkg/pagination"

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (p *PaginationInput) ToPagination() pagination.Pagination {
	return pagination.Pagination{
		Limit: p.Limit,
		Page:  p.Page,
	}
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
