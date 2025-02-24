package resolver

import (
	"context"

	"github.com/vshakitskiy/reddit_comments/internal/dataloader"
	"github.com/vshakitskiy/reddit_comments/internal/model"
)

type postResolver struct{ *Resolver }

func (r *postResolver) User(ctx context.Context, obj *model.Post) (*model.User, error) {
	return dataloader.GetUser(ctx, obj.UserID)
}
