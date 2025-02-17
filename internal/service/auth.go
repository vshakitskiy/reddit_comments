package service

import (
	"context"

	"github.com/vshakitskiy/reddit_comments/internal/model"
)

// TODO: implement all auth methods

func (s *Service) Register(
	ctx context.Context,
	username, password string,
) (*model.AuthPayload, error) {
	panic("not implemented")
}

func (s *Service) Login(
	ctx context.Context,
	username, password string,
) (*model.AuthPayload, error) {
	panic("not implemented")
}
