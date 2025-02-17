package inmemory

import (
	"context"

	"github.com/vshakitskiy/reddit_comments/internal/model"
)

// TODO: implement all post methods

func (r *InMemoryRepository) CreateComment(ctx context.Context, comment model.CommentMemory) (*model.CommentMemory, error) {
	panic("not implemented")
}

func (r *InMemoryRepository) GetCommentsByPostID(ctx context.Context, postID string, limit, offset int) ([]*model.CommentMemory, error) {
	panic("not implemented")
}

func (r *InMemoryRepository) GetReplies(ctx context.Context, commentID string, limit, offset int) ([]*model.CommentMemory, error) {
	panic("not implemented")
}
