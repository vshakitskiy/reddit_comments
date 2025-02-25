package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/vshakitskiy/reddit_comments/internal/model"
)

func (s *Service) CreateComment(
	input model.CommentInput,
	userID string,
) (*model.Comment, error) {
	post, _ := s.repo.GetPostByID(input.PostID)
	if post == nil {
		return nil, errors.New("post is not found")
	}

	comment := &model.Comment{
		ID:        uuid.New().String(),
		Content:   input.Content,
		UserID:    userID,
		ParentID:  input.ParentID,
		PostID:    input.PostID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.repo.CreateComment(comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *Service) GetComments(
	paginationInput model.PaginationInput,
	postID string,
) (*model.CommentsConnection, error) {
	pagination := paginationInput.ToPagination()
	return s.repo.GetComments(pagination, postID)
}

func (s *Service) GetReplies(
	paginationInput model.PaginationInput,
	parentID string,
) (*model.CommentsConnection, error) {
	pagination := paginationInput.ToPagination()
	return s.repo.GetReplies(pagination, parentID)
}
