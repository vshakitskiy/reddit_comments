package resolver

import (
	"context"

	"github.com/vshakitskiy/reddit_comments/internal/model"
)

type MutationResolver struct{ *Resolver }

// TODO: implement all mutation methods

func (r *MutationResolver) Register(ctx context.Context, cred model.AuthInput) (*model.AuthPayload, error) {
	panic("not implemented")
}

func (r *MutationResolver) Login(ctx context.Context, cred model.AuthInput) (*model.AuthPayload, error) {
	panic("not implemented")
}

func (r *MutationResolver) CreatePost(ctx context.Context, title string, description string, commentsEnabled bool, userID string) (*model.Post, error) {
	panic("not implemented")
}

func (r *MutationResolver) CreateComment(ctx context.Context, postID string, parentID *string, content string) (*model.Comment, error) {
	panic("not implemented")
}
