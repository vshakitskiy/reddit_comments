package inmemory

import (
	"context"

	"github.com/vshakitskiy/reddit_comments/internal/model"
)

// TODO: implement all user methods

func (r *InMemoryRepository) CreateUser(ctx context.Context, user model.UserMemory) (*model.UserMemory, error) {
	panic("not implemented")
}

func (r *InMemoryRepository) GetUserByID(ctx context.Context, id string) (*model.UserMemory, error) {
	panic("not implemented")
}
