package service

import "github.com/vshakitskiy/reddit_comments/internal/repository"

type Service struct {
	repo *repository.Repository
}

func NewService(
	repo *repository.Repository,
) *Service {
	return &Service{
		repo: repo,
	}
}
