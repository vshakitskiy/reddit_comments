package service

import (
	"context"

	"github.com/vshakitskiy/reddit_comments/internal/model"
)

// TODO: implement all post methods

func (s *Service) CreatePost(
	ctx context.Context,
	title, description string,
	commentsEnabled bool,
	userID string,
) (*model.Post, error) {
	panic("not implemented")
}

func (s *Service) GetPostByID(
	ctx context.Context,
	id string,
) (*model.Post, error) {
	panic("not implemented")
}

func (s *Service) GetAllPosts(
	ctx context.Context,
) ([]*model.Post, error) {
	panic("not implemented")
}
