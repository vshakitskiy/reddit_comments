package resolver

import (
	"context"

	"github.com/vshakitskiy/reddit_comments/internal/dataloader"
	"github.com/vshakitskiy/reddit_comments/internal/graph"
	"github.com/vshakitskiy/reddit_comments/internal/model"
)

type postResolver struct{ *Resolver }

func (r *Resolver) Post() graph.PostResolver {
	return &postResolver{r}
}

func (r *postResolver) User(ctx context.Context, obj *model.Post) (*model.User, error) {
	return dataloader.GetUser(ctx, obj.UserID)
}

func (r *postResolver) Comments(ctx context.Context, obj *model.Post, pagination model.PaginationInput) (*model.CommentsConnection, error) {
	if !obj.CommentsEnabled {
		return nil, nil
	}

	return r.svc.GetComments(pagination, obj.ID)
}
