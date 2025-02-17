package repository

import (
	"context"

	"github.com/vshakitskiy/reddit_comments/internal/model"
	"github.com/vshakitskiy/reddit_comments/internal/repository/inmemory"
)

type Repository interface {
	CreatePost(
		ctx context.Context,
		post model.Post,
	) (*model.Post, error)
	GetPostByID(
		ctx context.Context,
		id string,
	) (*model.Post, error)
	GetAllPosts(
		ctx context.Context,
	) ([]*model.Post, error)

	CreateComment(
		ctx context.Context,
		comment model.Comment,
	) (*model.Comment, error)
	GetCommentsByPostID(
		ctx context.Context,
		postID string,
		limit, offset int,
	) ([]*model.Comment, error)
	GetReplies(
		ctx context.Context,
		commentID string,
		limit, offset int,
	) ([]*model.Comment, error)
}

func NewRepository(opt string) Repository {
	switch opt {
	case "inmemory":
		return inmemory.NewInMemoryRepository()
	case "postgres":
		panic("not implemented")
	default:
		panic("invalid repository option")
	}
}
