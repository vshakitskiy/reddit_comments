package dataloader

import (
	"context"
	"time"

	"github.com/vikstrous/dataloadgen"
	"github.com/vshakitskiy/reddit_comments/internal/model"
	"github.com/vshakitskiy/reddit_comments/internal/repository"
)

func NewUserLoader(repo *repository.Repository) *dataloadgen.Loader[string, *model.User] {
	return dataloadgen.NewLoader(
		repo.GetUsersByIDs,
		dataloadgen.WithWait(time.Millisecond),
	)
}

func GetUser(ctx context.Context, id string) (*model.User, error) {
	loader := LoaderFromCtx(ctx)
	return loader.UserLoader.Load(ctx, id)
}

func GetUsers(ctx context.Context, ids []string) ([]*model.User, error) {
	loader := LoaderFromCtx(ctx)
	return loader.UserLoader.LoadAll(ctx, ids)
}
