package resolver

import (
	"context"

	"github.com/vshakitskiy/reddit_comments/internal/model"
)

type MutationResolver struct{ *Resolver }

func (r *MutationResolver) CreatePost(
	ctx context.Context,
	title string,
	description string,
	commentsEnabled bool,
	userID string,
) (*model.Post, error) {
	return r.service.CreatePost(
		ctx,
		title,
		description,
		commentsEnabled,
		userID,
	)
}

// CreateComment is the resolver for the createComment field.
func (r *MutationResolver) CreateComment(ctx context.Context, postID string, parentID *string, content string) (*model.Comment, error) {
	panic("not implemented")
}
