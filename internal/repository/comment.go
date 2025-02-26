package repository

import (
	"context"
	"errors"

	"github.com/vshakitskiy/reddit_comments/internal/model"
	"github.com/vshakitskiy/reddit_comments/internal/repository/pg"
	"github.com/vshakitskiy/reddit_comments/pkg/pagination"
	"gorm.io/gorm"
)

func (r *Repository) CreateComment(comment *model.Comment) error {
	err := r.db.Create(comment).Error
	if err != nil {
		return errors.New("unable to create comment")
	}

	return nil
}

func (r *Repository) GetCommentByID(commentID string) (*model.Comment, error) {
	var comment *model.Comment

	tx := r.db.First(&comment, "id = ?", commentID)
	if tx.Error != nil {
		return nil, errors.New("comment is not found")
	}

	return comment, nil
}

func (r *Repository) GetCommentsByIDs(
	_ context.Context,
	ids []string,
) ([]*model.Comment, []error) {
	var comments []*model.Comment

	tx := r.db.Where("id IN ?", ids).Find(&comments)
	if tx.Error != nil {
		return nil, []error{tx.Error}
	}

	comments, err := onIdOrder(ids, comments)
	if err != nil {
		return nil, []error{err}
	}

	return comments, nil
}

func (r *Repository) GetReplies(
	pagination pagination.Pagination,
	parentID string,
	totalReplies int64,
) (*model.CommentsConnection, error) {
	var comments []*model.Comment

	tx := r.db.Scopes(pg.Paginate(
		comments,
		&pagination,
		r.db,
		totalReplies,
	)).Where("parent_id = ?", parentID).Find(&comments)

	if tx.Error != nil {
		return nil, errors.New("unable to get replies")
	}

	meta := model.PaginationToConnectionMeta(pagination)

	return &model.CommentsConnection{
		Rows: comments,
		Meta: meta,
	}, nil
}

func (r *Repository) GetComments(
	pagination pagination.Pagination,
	postID string,
	totalComments int64,
) (*model.CommentsConnection, error) {
	var comments []*model.Comment

	tx := r.db.Scopes(pg.Paginate(
		comments,
		&pagination,
		r.db,
		totalComments,
	)).Where("post_id = ?", postID).Where("parent_id IS NULL").Find(&comments)

	if tx.Error != nil {
		return nil, errors.New("unable to get comments")
	}

	meta := model.PaginationToConnectionMeta(pagination)

	return &model.CommentsConnection{
		Rows: comments,
		Meta: meta,
	}, nil
}

func (r *Repository) CountComments(postID string, totalRows *int64) {
	r.db.Model(&model.Comment{}).Where("post_id = ?", postID).Where("parent_id IS NULL").Count(totalRows)
}

func (r *Repository) CountReplies(parentID string, totalRows *int64) {
	r.db.Model(&model.Comment{}).Where("parent_id = ?", parentID).Count(totalRows)
}

func (r *Repository) IncrementRepliesCount(parentID string) error {
	tx := r.db.Model(&model.Comment{}).Where("id = ?", parentID).Update("total_replies", gorm.Expr("total_replies + 1"))
	if tx.Error != nil {
		return errors.New("unable to increment replies count")
	}
	return nil
}
