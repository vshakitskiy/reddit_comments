package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"
	"fmt"

	"github.com/vshakitskiy/reddit_comments/internal/model"
)

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, title string, description string, commentsEnabled bool) (*model.Post, error) {
	panic(fmt.Errorf("not implemented: CreatePost - createPost"))
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, postID string, parentID *string, content string) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented: CreateComment - createComment"))
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented: Posts - posts"))
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	panic(fmt.Errorf("not implemented: Post - post"))
}

// NewComment is the resolver for the newComment field.
func (r *subscriptionResolver) NewComment(ctx context.Context, postID string) (<-chan *model.Comment, error) {
	panic(fmt.Errorf("not implemented: NewComment - newComment"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
