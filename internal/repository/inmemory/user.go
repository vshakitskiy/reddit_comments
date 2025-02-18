package inmemory

import (
	"context"
	"errors"

	"github.com/vshakitskiy/reddit_comments/internal/model"
)

// TODO: implement all user methods

func (r *InMemoryRepository) CreateUser(ctx context.Context, user model.UserMemory) (*model.UserMemory, error) {
	r.users[user.ID] = &user
	return &user, nil
}

func (r *InMemoryRepository) GetUserByID(ctx context.Context, id string) (*model.UserMemory, error) {
	panic("not implemented")
}

func (r *InMemoryRepository) GetUserByUsername(
	ctx context.Context,
	username string,
) (*model.UserMemory, error) {
	for _, user := range r.users {
		if user.Username == username {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}
