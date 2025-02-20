package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/vshakitskiy/reddit_comments/internal/model"
	"github.com/vshakitskiy/reddit_comments/pkg/bcrypt"
	"github.com/vshakitskiy/reddit_comments/pkg/jwt"
)

func (s *Service) Register(
	credentials model.RegisterInput,
) (*model.AuthPayload, error) {
	if u, _ := s.repo.GetUserByEmail(
		credentials.Email,
	); u != nil {
		return nil, errors.New("user is already exists")
	}

	passwordHash := bcrypt.Hash(credentials.Password)

	user := &model.User{
		Id:           uuid.New().String(),
		Username:     credentials.Username,
		Email:        credentials.Email,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}

	token, err := jwt.Generate(user.Id)
	if err != nil {
		return nil, err
	}

	return &model.AuthPayload{
		Token: token,
		User:  user,
	}, nil
}

func (s *Service) Login(
	credentials model.LoginInput,
) (*model.AuthPayload, error) {
	user, err := s.repo.GetUserByEmail(
		credentials.Email,
	)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.Compare(user.PasswordHash, credentials.Password); err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := jwt.Generate(user.Id)
	if err != nil {
		return nil, err
	}

	return &model.AuthPayload{
		Token: token,
		User:  user,
	}, nil
}
