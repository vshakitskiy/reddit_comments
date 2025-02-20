package resolver

// THIS CODE WILL BE UPDATED WITH SCHEMA CHANGES. PREVIOUS IMPLEMENTATION FOR SCHEMA CHANGES WILL BE KEPT IN THE COMMENT SECTION. IMPLEMENTATION FOR UNCHANGED SCHEMA WILL BE KEPT.

import (
	"context"

	"github.com/vshakitskiy/reddit_comments/internal/graph"
	"github.com/vshakitskiy/reddit_comments/internal/model"
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

// User is the resolver for the user field.
func (r *commentResolver) User(ctx context.Context, obj *model.Comment) (*model.User, error) {
	panic("not implemented")
}

// Parent is the resolver for the parent field.
func (r *commentResolver) Parent(ctx context.Context, obj *model.Comment) (*model.Comment, error) {
	panic("not implemented")
}

// Replies is the resolver for the replies field.
func (r *commentResolver) Replies(ctx context.Context, obj *model.Comment, pagination model.PaginationInput) (*model.CommentsConnection, error) {
	panic("not implemented")
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.PostInput) (*model.Post, error) {
	panic("not implemented")
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, input model.CommentInput) (*model.Comment, error) {
	panic("not implemented")
}

// TotalComments is the resolver for the totalComments field.
func (r *postResolver) TotalComments(ctx context.Context, obj *model.Post) (int32, error) {
	panic("not implemented")
}

// User is the resolver for the user field.
func (r *postResolver) User(ctx context.Context, obj *model.Post) (*model.User, error) {
	panic("not implemented")
}

// Comments is the resolver for the comments field.
func (r *postResolver) Comments(ctx context.Context, obj *model.Post, pagination model.PaginationInput) (*model.CommentsConnection, error) {
	panic("not implemented")
}

// GetPosts is the resolver for the getPosts field.
func (r *queryResolver) GetPosts(ctx context.Context, pagination model.PaginationInput) (*model.PostsConnection, error) {
	panic("not implemented")
}

// GetPostByID is the resolver for the getPostById field.
func (r *queryResolver) GetPostByID(ctx context.Context, postID string) (*model.Post, error) {
	panic("not implemented")
}

// GetReplies is the resolver for the getReplies field.
func (r *queryResolver) GetReplies(ctx context.Context, commentID string) (*model.CommentsConnection, error) {
	panic("not implemented")
}

// NewComment is the resolver for the newComment field.
func (r *subscriptionResolver) NewComment(ctx context.Context, postID string) (<-chan *model.Comment, error) {
	panic("not implemented")
}

// Comment returns graph.CommentResolver implementation.
func (r *Resolver) Comment() graph.CommentResolver { return &commentResolver{r} }

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Post returns graph.PostResolver implementation.
func (r *Resolver) Post() graph.PostResolver { return &postResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

// Subscription returns graph.SubscriptionResolver implementation.
func (r *Resolver) Subscription() graph.SubscriptionResolver { return &subscriptionResolver{r} }

type commentResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	type Resolver struct{}
*/
