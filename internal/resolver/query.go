package resolver

import (
	"context"

	"github.com/vshakitskiy/reddit_comments/internal/graph"
	"github.com/vshakitskiy/reddit_comments/internal/model"
)

type queryResolver struct{ *Resolver }

func (r *Resolver) Query() graph.QueryResolver {
	return &queryResolver{r}
}

func (r *queryResolver) GetPosts(
	ctx context.Context,
	pagination model.PaginationInput,
) (*model.PostsConnection, error) {
	return r.svc.GetPosts(pagination)
}

func (r *queryResolver) GetPostByID(
	ctx context.Context,
	postID string,
) (*model.Post, error) {
	return r.svc.GetPostByID(postID)
}

func (r *queryResolver) GetReplies(ctx context.Context, commentID string, pagination model.PaginationInput) (*model.CommentsConnection, error) {
	panic("not implemented")

}
