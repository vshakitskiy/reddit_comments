package resolver

import (
	"context"

	"github.com/vshakitskiy/reddit_comments/internal/dataloader"
	"github.com/vshakitskiy/reddit_comments/internal/graph"
	"github.com/vshakitskiy/reddit_comments/internal/model"
)

type commentResolver struct{ *Resolver }

func (r *Resolver) Comment() graph.CommentResolver {
	return &commentResolver{r}
}

func (r *commentResolver) User(ctx context.Context, obj *model.Comment) (*model.User, error) {
	return dataloader.GetUser(ctx, obj.UserID)
}

func (r *commentResolver) Parent(ctx context.Context, obj *model.Comment) (*model.Comment, error) {
	if obj.ParentID == nil {
		return nil, nil
	}
	return dataloader.GetParent(ctx, *obj.ParentID)
}

func (r *commentResolver) Replies(ctx context.Context, obj *model.Comment, pagination model.PaginationInput) (*model.CommentsConnection, error) {
	return r.svc.GetReplies(pagination, obj.ID, &obj.TotalReplies)
}
