package service

import (
	"github.com/vshakitskiy/reddit_comments/internal/model"
)

func (s *Service) CreateComment(
	input model.CommentInput,
	userID string,
) (*model.Comment, error) {
	panic("not implemented")
}
