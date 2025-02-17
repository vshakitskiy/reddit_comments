package inmemory

import (
	"context"
	"errors"
	"sync"

	"github.com/vshakitskiy/reddit_comments/internal/model"
)

type InMemoryRepository struct {
	mu       sync.RWMutex
	posts    map[string]*model.Post
	comments map[string]*model.Comment
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		posts:    make(map[string]*model.Post),
		comments: make(map[string]*model.Comment),
	}
}

func (r *InMemoryRepository) CreatePost(
	ctx context.Context,
	post model.Post,
) (*model.Post, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.posts[post.ID]; exists {
		return nil, errors.New("post already exists")
	}

	r.posts[post.ID] = &post
	return &post, nil
}

// TODO: implement all inmemory methods

func (r *InMemoryRepository) GetPostByID(ctx context.Context, id string) (*model.Post, error) {
	panic("not implemented")
}

func (r *InMemoryRepository) GetAllPosts(ctx context.Context) ([]*model.Post, error) {
	panic("not implemented")
}

func (r *InMemoryRepository) CreateComment(ctx context.Context, comment model.Comment) (*model.Comment, error) {
	panic("not implemented")
}

func (r *InMemoryRepository) GetCommentsByPostID(ctx context.Context, postID string, limit, offset int) ([]*model.Comment, error) {
	panic("not implemented")
}

func (r *InMemoryRepository) GetReplies(ctx context.Context, commentID string, limit, offset int) ([]*model.Comment, error) {
	panic("not implemented")
}
