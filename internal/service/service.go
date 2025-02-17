package service

import (
	"github.com/vshakitskiy/reddit_comments/internal/repository"
)

type Service struct {
	repository repository.Repository
}

func NewService(
	repository repository.Repository,
) *Service {
	return &Service{
		repository: repository,
	}
}

// type ServiceImpl interface {
// 	CreatePost(
// 		ctx context.Context,
// 		title, description string,
// 		commentsEnabled bool,
// 		userID string,
// 	) (*model.Post, error)
// 	GetPostByID(
// 		ctx context.Context,
// 		id string,
// 	) (*model.Post, error)
// 	GetAllPosts(
// 		ctx context.Context,
// 	) ([]*model.Post, error)

// 	CreateComment(
// 		ctx context.Context,
// 		postID string,
// 		parentID *string,
// 		content string,
// 		userID string,
// 	) (*model.Comment, error)
// 	GetCommentsByPostID(
// 		ctx context.Context,
// 		postID string,
// 		limit, offset int,
// 	) ([]*model.Comment, error)
// 	GetReplies(
// 		ctx context.Context,
// 		commentID string,
// 		limit, offset int,
// 	) ([]*model.Comment, error)
// }
