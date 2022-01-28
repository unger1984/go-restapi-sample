package users

import (
	domain "awcoding.com/back/domain/users"
)

type User struct {
	Id       int    `json:"id" db:"id"`
	Email    string `json:"email" db:"email" binding:"required"`
	Password string `json:"-" db:"password"`
}

func (u User) toEntity() *domain.User {
	return &domain.User{Id: u.Id, Email: u.Email, Password: u.Password}
}

func fromEntity(u *domain.User) *User {
	return &User{Id: u.Id, Email: u.Email, Password: u.Password}
}
