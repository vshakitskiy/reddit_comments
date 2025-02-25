package repository

import (
	"errors"

	"github.com/vshakitskiy/reddit_comments/internal/model"
	"github.com/vshakitskiy/reddit_comments/internal/repository/pg"
	"github.com/vshakitskiy/reddit_comments/pkg/pagination"
	"gorm.io/gorm"
)

func (r *Repository) GetPosts(
	pagination pagination.Pagination,
) (*model.PostsConnection, error) {
	var posts []*model.Post

	var totalRows int64
	r.db.Model(&model.Post{}).Count(&totalRows)

	tx := r.db.Scopes(pg.Paginate(
		posts,
		&pagination,
		r.db,
		totalRows,
	)).Find(&posts)

	if tx.Error != nil {
		return nil, errors.New("unable to get posts")
	}

	meta := model.PaginationToConnectionMeta(pagination)

	return &model.PostsConnection{
		Rows: posts,
		Meta: meta,
	}, nil
}

func (r *Repository) GetPostByID(postID string) (*model.Post, error) {
	var post model.Post

	tx := r.db.First(&post, "id = ?", postID)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return nil, errors.New("post is not found")
		}
		return nil, tx.Error
	}

	return &post, nil
}

func (r *Repository) CreatePost(
	post *model.Post,
) error {
	err := r.db.Create(post).Error
	if err != nil {
		return errors.New("unable to create post")
	}

	return nil
}
