package resolver

import (
	"context"

	"github.com/vshakitskiy/reddit_comments/internal/model"
)

type SubscriptionResolver struct{ *Resolver }

// TODO: implement subscription method
func (r *SubscriptionResolver) NewComment(ctx context.Context, postID string) (<-chan *model.Comment, error) {
	panic("not implemented")
}
