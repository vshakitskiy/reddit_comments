package resolver

import (
	"context"

	"github.com/vshakitskiy/reddit_comments/internal/model"
)

type QueryResolver struct{ *Resolver }

// TODO: implement all query methods

func (r *QueryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	panic("not implemented")
}

func (r *QueryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	panic("not implemented")
}

func (r *QueryResolver) Comments(ctx context.Context, postID string) ([]*model.Comment, error) {
	panic("not implemented")
}

func (r *QueryResolver) Replies(ctx context.Context, commentID string) ([]*model.Comment, error) {
	panic("not implemented")
}
