package resolver

import (
	"context"

	"github.com/vshakitskiy/reddit_comments/internal/model"
)

type mutationResolver struct{ *Resolver }

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
