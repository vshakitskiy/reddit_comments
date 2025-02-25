package resolver

// THIS CODE WILL BE UPDATED WITH SCHEMA CHANGES. PREVIOUS IMPLEMENTATION FOR SCHEMA CHANGES WILL BE KEPT IN THE COMMENT SECTION. IMPLEMENTATION FOR UNCHANGED SCHEMA WILL BE KEPT.

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

// // TotalRow is the resolver for the totalRow field.
// func (r *connectionMetaResolver) TotalRow(ctx context.Context, obj *model.ConnectionMeta) (int32, error) {
// 	panic("not implemented")
// }

// // TotalPages is the resolver for the totalPages field.
// func (r *connectionMetaResolver) TotalPages(ctx context.Context, obj *model.ConnectionMeta) (int32, error) {
// 	panic("not implemented")
// }

// // Limit is the resolver for the limit field.
// func (r *connectionMetaResolver) Limit(ctx context.Context, obj *model.ConnectionMeta) (int32, error) {
// 	panic("not implemented")
// }

// // Page is the resolver for the page field.
// func (r *connectionMetaResolver) Page(ctx context.Context, obj *model.ConnectionMeta) (int32, error) {
// 	panic("not implemented")
// }

// // TotalComments is the resolver for the totalComments field.
// func (r *postResolver) TotalComments(ctx context.Context, obj *model.Post) (int32, error) {
// 	panic("not implemented")
// }

// // Comments is the resolver for the comments field.
// func (r *postResolver) Comments(ctx context.Context, obj *model.Post, pagination model.PaginationInput) (*model.CommentsConnection, error) {
// 	panic("not implemented")
// }

// // ConnectionMeta returns graph.ConnectionMetaResolver implementation.
// func (r *Resolver) ConnectionMeta() graph.ConnectionMetaResolver { return &connectionMetaResolver{r} }

// type connectionMetaResolver struct{ *Resolver }
