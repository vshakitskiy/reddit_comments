package resolver

import (
	"context"

	"github.com/vshakitskiy/reddit_comments/internal/model"
)

type QueryResolver struct{ *Resolver }

// Posts is the resolver for the posts field.
func (r *QueryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	// GET ALL POSTS
	panic("not implemented")
}

// Post is the resolver for the post field.
func (r *QueryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	// GET POST
	panic("not implemented")
}

// Comments is the resolver for the comments field.
func (r *QueryResolver) Comments(ctx context.Context, postID string) ([]*model.Comment, error) {
	panic("not implemented")
}

// Replies is the resolver for the replies field.
func (r *QueryResolver) Replies(ctx context.Context, commentID string) ([]*model.Comment, error) {
	panic("not implemented")
}
