package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/vshakitskiy/reddit_comments/internal/model"
)

func (s *Service) GetPosts(
	paginationInput model.PaginationInput,
) (*model.PostsConnection, error) {
	pagination := paginationInput.ToPagination()
	return s.repo.GetPosts(pagination)
}

func (s *Service) GetPostByID(
	postID string,
) (*model.Post, error) {
	return s.repo.GetPostByID(postID)
}

func (s *Service) CreatePost(
	input model.PostInput,
	id string,
) (*model.Post, error) {
	post := &model.Post{
		ID:              uuid.New().String(),
		Title:           input.Title,
		Description:     input.Description,
		CommentsEnabled: input.CommentsEnabled,
		TotalComments:   0,
		UserID:          id,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	err := s.repo.CreatePost(post)
	if err != nil {
		return nil, err
	}

	return post, nil
}
