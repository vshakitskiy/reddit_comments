package service

import (
	"context"

	"github.com/vshakitskiy/reddit_comments/internal/model"
)

// TODO: implement all comment methods

func (s *Service) CreateComment(
	ctx context.Context,
	postID string,
	parentID *string,
	content string,
	userID string,
) (*model.Comment, error) {
	panic("not implemented")
}

func (s *Service) GetCommentsByPostID(
	ctx context.Context,
	postID string,
	limit, offset int,
) ([]*model.Comment, error) {
	panic("not implemented")
}

func (s *Service) GetReplies(
	ctx context.Context,
	commentID string,
	limit, offset int,
) ([]*model.Comment, error) {
	panic("not implemented")
}
