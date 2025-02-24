package model

import "time"

type User struct {
	ID           string `gorm:"primaryKey" json:"id"`
	Username     string `gorm:"not null" json:"username"`
	Email        string `gorm:"unique;not null" json:"email"`
	PasswordHash string `gorm:"not null" json:"password_hash"`

	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}

func (u *User) GetID() string {
	return u.ID
}

type AuthPayload struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}
