package inmemory

import (
	"context"
	"errors"

	"github.com/vshakitskiy/reddit_comments/internal/model"
)

// TODO: implement all post methods

func (r *InMemoryRepository) CreatePost(
	ctx context.Context,
	post model.PostMemory,
) (*model.PostMemory, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.posts[post.ID]; exists {
		return nil, errors.New("post already exists")
	}

	r.posts[post.ID] = &post
	return &post, nil
}

func (r *InMemoryRepository) GetPostByID(ctx context.Context, id string) (*model.PostMemory, error) {
	panic("not implemented")
}

func (r *InMemoryRepository) GetAllPosts(ctx context.Context) ([]*model.PostMemory, error) {
	panic("not implemented")
}
