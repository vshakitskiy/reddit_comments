package repository

import (
	"context"

	"github.com/vshakitskiy/reddit_comments/internal/model"
	"github.com/vshakitskiy/reddit_comments/internal/repository/inmemory"
)

type Repository interface {
	CreatePost(
		ctx context.Context,
		post model.PostMemory,
	) (*model.PostMemory, error)
	GetPostByID(
		ctx context.Context,
		id string,
	) (*model.PostMemory, error)
	GetAllPosts(
		ctx context.Context,
	) ([]*model.PostMemory, error)

	CreateComment(
		ctx context.Context,
		comment model.CommentMemory,
	) (*model.CommentMemory, error)
	GetCommentsByPostID(
		ctx context.Context,
		postID string,
		limit, offset int,
	) ([]*model.CommentMemory, error)
	GetReplies(
		ctx context.Context,
		commentID string,
		limit, offset int,
	) ([]*model.CommentMemory, error)

	CreateUser(
		ctx context.Context,
		user model.UserMemory,
	) (*model.UserMemory, error)
	GetUserByID(
		ctx context.Context,
		id string,
	) (*model.UserMemory, error)
	GetUserByUsername(
		ctx context.Context,
		username string,
	) (*model.UserMemory, error)
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
