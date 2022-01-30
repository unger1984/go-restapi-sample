package models

import (
	"awcoding.com/back/pkg/domain/entities"
)

type User struct {
	Id       int     `db:"id"`
	Email    string  `db:"email" binding:"required"`
	AvatarId *int    `db:"avatarId" `
	Password string  `db:"password"`
	Avatar   *Upload `db:"avatar"`
}

func (u User) ToEntity() *entities.User {
	return &entities.User{Id: u.Id, Email: u.Email, Password: u.Password, AvatarId: u.AvatarId, Avatar: u.Avatar.ToEntity()}
}
