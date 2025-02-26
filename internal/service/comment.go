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
	post, err := s.repo.GetPostByID(input.PostID)
	if err != nil {
		return nil, err
	}

	if !post.CommentsEnabled {
		return nil, errors.New("comments are disabled")
	}

	if input.ParentID != nil {
		_, err := s.repo.GetCommentByID(*input.ParentID)
		if err != nil {
			return nil, err
		}
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

	err = s.repo.CreateComment(comment)
	if err != nil {
		return nil, err
	}

	err = s.repo.IncrementCommentCount(input.PostID)
	if err != nil {
		return comment, err
	}

	if input.ParentID != nil {
		err = s.repo.IncrementRepliesCount(*input.ParentID)
		if err != nil {
			return comment, err
		}
	}

	return comment, nil
}

func (s *Service) GetComments(
	paginationInput model.PaginationInput,
	postID string,
) (*model.CommentsConnection, error) {
	post, err := s.repo.GetPostByID(postID)
	if err != nil {
		return nil, err
	}

	if !post.CommentsEnabled {
		return nil, nil
	}

	pagination := paginationInput.ToPagination()
	var total int64
	s.repo.CountComments(postID, &total)

	return s.repo.GetComments(pagination, postID, total)
}

func (s *Service) GetReplies(
	paginationInput model.PaginationInput,
	parentID string,
	totalReplies *int32,
) (*model.CommentsConnection, error) {
	pagination := paginationInput.ToPagination()
	var total int64

	if totalReplies != nil {
		total = int64(*totalReplies)
	} else {
		comment, err := s.repo.GetCommentByID(parentID)
		if err != nil {
			return nil, err
		}
		total = int64(comment.TotalReplies)
	}

	return s.repo.GetReplies(pagination, parentID, total)
}
