package service

import (
	"context"
	"errors"

	"github.com/vshakitskiy/reddit_comments/internal/model"
	"github.com/vshakitskiy/reddit_comments/pkg/bcrypt"
	"github.com/vshakitskiy/reddit_comments/pkg/jwt"
)

// TODO: implement all auth methods

func (s *Service) Register(
	ctx context.Context,
	username, password string,
) (*model.AuthPayload, error) {
	createdUser, err := s.UserCreate(ctx, username, password)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Generate(createdUser.ID)
	if err != nil {
		return nil, err
	}

	return &model.AuthPayload{
		Token: token,
		User:  createdUser,
	}, nil
}

func (s *Service) Login(
	ctx context.Context,
	username, password string,
) (*model.AuthPayload, error) {
	dbUser, err := s.repository.GetUserByUsername(
		ctx,
		username,
	)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.Compare(dbUser.PasswordHash, password); err != nil {
		return nil, errors.New("invalid password")
	}

	token, err := jwt.Generate(dbUser.ID)
	if err != nil {
		return nil, err
	}

	return &model.AuthPayload{
		Token: token,
		User:  (*model.User)(dbUser),
	}, nil
}
