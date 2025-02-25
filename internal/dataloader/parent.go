package dataloader

import (
	"context"
	"time"

	"github.com/vikstrous/dataloadgen"
	"github.com/vshakitskiy/reddit_comments/internal/model"
	"github.com/vshakitskiy/reddit_comments/internal/repository"
)

func NewParentLoader(repo *repository.Repository) *dataloadgen.Loader[string, *model.Comment] {
	return dataloadgen.NewLoader(
		repo.GetCommentsByIDs,
		dataloadgen.WithWait(time.Millisecond),
	)
}

func GetParent(ctx context.Context, id string) (*model.Comment, error) {
	loader := LoaderFromCtx(ctx)
	return loader.ParentLoader.Load(ctx, id)
}

func GetParents(ctx context.Context, ids []string) ([]*model.Comment, error) {
	loader := LoaderFromCtx(ctx)
	return loader.ParentLoader.LoadAll(ctx, ids)
}
