package inmemory

import (
	"sync"

	"github.com/vshakitskiy/reddit_comments/internal/model"
)

type InMemoryRepository struct {
	mu       sync.RWMutex
	posts    map[string]*model.PostMemory
	comments map[string]*model.CommentMemory
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		posts:    make(map[string]*model.PostMemory),
		comments: make(map[string]*model.CommentMemory),
	}
}
