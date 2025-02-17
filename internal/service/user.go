package service

import (
	"context"

	"github.com/vshakitskiy/reddit_comments/internal/model"
)

// TODO: implement all user methods

func UserCreate(
	ctx context.Context,
	username, password string,
) (*model.User, error) {
	panic("not implemented")
}

func GetUserByID(
	ctx context.Context,
	id string,
) (*model.User, error) {
	panic("not implemented")
}
