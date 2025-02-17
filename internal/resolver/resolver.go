package resolver

// THIS CODE WILL BE UPDATED WITH SCHEMA CHANGES. PREVIOUS IMPLEMENTATION FOR SCHEMA CHANGES WILL BE KEPT IN THE COMMENT SECTION. IMPLEMENTATION FOR UNCHANGED SCHEMA WILL BE KEPT.

import (
	"github.com/vshakitskiy/reddit_comments/internal/graph"
	"github.com/vshakitskiy/reddit_comments/internal/service"
)

type Resolver struct {
	service *service.Service
}

func NewResolver(service *service.Service) *Resolver {
	return &Resolver{service: service}
}

func (r *Resolver) Mutation() graph.MutationResolver {
	return &MutationResolver{r}
}

func (r *Resolver) Query() graph.QueryResolver {
	return &QueryResolver{r}
}

func (r *Resolver) Subscription() graph.SubscriptionResolver {
	return &SubscriptionResolver{r}
}
