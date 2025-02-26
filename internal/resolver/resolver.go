package resolver

import (
	"github.com/vshakitskiy/reddit_comments/internal/service"
)

type Resolver struct {
	svc *service.Service
}

func NewResolver(svc *service.Service) *Resolver {
	return &Resolver{
		svc: svc,
	}
}
