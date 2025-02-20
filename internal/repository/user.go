package repository

import (
	"errors"

	"github.com/vshakitskiy/reddit_comments/internal/model"
)

func (r *Repository) CreateUser(
	user *model.User,
) error {
	if err := r.db.Create(user).Error; err != nil {
		return errors.New("unable to create user")
	}
	return nil
}

func (r *Repository) GetUserByEmail(
	email string,
) (*model.User, error) {
	var user model.User

	tx := r.db.Where("email = ?", email).First(&user)
	if tx.Error != nil {
		return nil, errors.New("user is not found")
	}

	return &user, nil
}
