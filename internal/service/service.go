package service

import (
	"github.com/vshakitskiy/reddit_comments/internal/pubsub"
	"github.com/vshakitskiy/reddit_comments/internal/repository"
)

type Service struct {
	repo          *repository.Repository
	commentPubSub *pubsub.CommentPubSub
}

func NewService(
	repo *repository.Repository,
	commentPubSub *pubsub.CommentPubSub,
) *Service {
	return &Service{
		repo:          repo,
		commentPubSub: commentPubSub,
	}
}
