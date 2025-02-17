package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/vshakitskiy/reddit_comments/internal/model"
	"github.com/vshakitskiy/reddit_comments/internal/repository"
)

type ServiceImpl interface {
	CreatePost(
		ctx context.Context,
		title, description string,
		commentsEnabled bool,
		userID string,
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
		postID string,
		parentID *string,
		content string,
		userID string,
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

type Service struct {
	ServiceImpl
	repository repository.Repository
}

func NewService(
	repository repository.Repository,
) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) CreatePost(
	ctx context.Context,
	title, description string,
	commentsEnabled bool,
	userID string,
) (*model.Post, error) {
	post := model.Post{
		ID:              uuid.New().String(),
		Title:           title,
		Description:     description,
		CommentsEnabled: commentsEnabled,
		User:            nil,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		CommentCount:    0,
		Comments:        nil,
	}

	if post, err := s.repository.CreatePost(ctx, post); err != nil {
		return nil, err
	} else {
		return post, nil
	}
}

func (s *Service) GetPostByID(
	ctx context.Context,
	id string,
) (*model.Post, error) {
	panic("not implemented")
}

func (s *Service) GetAllPosts(
	ctx context.Context,
) ([]*model.Post, error) {
	panic("not implemented")
}

func (s *Service) CreateComment(
	ctx context.Context,
	postID string,
	parentID *string,
	content string,
	userID string,
) (*model.Comment, error) {
	panic("not implemented")
}

func (s *Service) GetCommentsByPostID(
	ctx context.Context,
	postID string,
	limit, offset int,
) ([]*model.Comment, error) {
	panic("not implemented")
}

func (s *Service) GetReplies(
	ctx context.Context,
	commentID string,
	limit, offset int,
) ([]*model.Comment, error) {
	panic("not implemented")
}
