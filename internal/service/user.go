package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/vshakitskiy/reddit_comments/internal/model"
	"github.com/vshakitskiy/reddit_comments/pkg/bcrypt"
)

// TODO: implement all user methods

func (s *Service) UserCreate(
	ctx context.Context,
	username, password string,
) (*model.User, error) {
	existingUser, err := s.repository.GetUserByUsername(ctx, username)
	if existingUser != nil {
		return nil, fmt.Errorf("user with username %s already exists", username)
	}

	passwordHash := bcrypt.Hash(password)

	user := model.UserMemory{
		ID:           uuid.New().String(),
		Username:     username,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	createdUser, err := s.repository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return (*model.User)(createdUser), nil
}

func (s *Service) GetUserByID(
	ctx context.Context,
	id string,
) (*model.User, error) {
	panic("not implemented")
}

func (s *Service) GetUserByUsername(
	ctx context.Context,
	username string,
) (*model.User, error) {
	panic("not implemented")
}
