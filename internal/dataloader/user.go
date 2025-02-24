package dataloader

import (
	"context"

	"github.com/vshakitskiy/reddit_comments/internal/model"
)

func GetUser(ctx context.Context, id string) (*model.User, error) {
	loader := LoaderFromCtx(ctx)
	return loader.UserLoader.Load(ctx, id)
}

func GetUsers(ctx context.Context, ids []string) ([]*model.User, error) {
	loader := LoaderFromCtx(ctx)
	return loader.UserLoader.LoadAll(ctx, ids)
}
