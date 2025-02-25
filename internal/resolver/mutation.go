package resolver

import (
	"context"

	"github.com/vshakitskiy/reddit_comments/internal/api/middleware"
	"github.com/vshakitskiy/reddit_comments/internal/graph"
	"github.com/vshakitskiy/reddit_comments/internal/model"
)

type mutationResolver struct{ *Resolver }

func (r *Resolver) Mutation() graph.MutationResolver {
	return &mutationResolver{r}
}

func (r *mutationResolver) Register(
	ctx context.Context,
	credentials model.RegisterInput,
) (*model.AuthPayload, error) {
	return r.svc.Register(credentials)
}

func (r *mutationResolver) Login(
	ctx context.Context,
	credentials model.LoginInput,
) (*model.AuthPayload, error) {
	return r.svc.Login(credentials)
}

func (r *mutationResolver) CreatePost(
	ctx context.Context,
	input model.PostInput,
) (*model.Post, error) {
	id := middleware.ExtractUserID(ctx)
	return r.svc.CreatePost(input, id)
}

func (r *mutationResolver) CreateComment(
	ctx context.Context,
	input model.CommentInput,
) (*model.Comment, error) {
	id := middleware.ExtractUserID(ctx)
	return r.svc.CreateComment(input, id)
}
