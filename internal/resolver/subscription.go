package resolver

import (
	"context"

	"github.com/vshakitskiy/reddit_comments/internal/graph"
	"github.com/vshakitskiy/reddit_comments/internal/model"
)

type subscriptionResolver struct{ *Resolver }

func (r *Resolver) Subscription() graph.SubscriptionResolver {
	return &subscriptionResolver{r}
}

func (r *subscriptionResolver) NewComment(ctx context.Context, postID string) (<-chan *model.Comment, error) {
	panic("not implemented")
}
