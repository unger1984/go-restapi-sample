package users

import (
	"awcoding.com/back/src/data/uploads"
	domain "awcoding.com/back/src/domain/users"
)

type User struct {
	Id       int             `db:"id"`
	Email    string          `db:"email" binding:"required"`
	AvatarId *int            `db:"avatarId" `
	Password string          `db:"password"`
	Avatar   *uploads.Upload `db:"avatar"`
}

func (u User) ToEntity() *domain.User {
	return &domain.User{Id: u.Id, Email: u.Email, Password: u.Password, AvatarId: u.AvatarId, Avatar: u.Avatar.ToEntity()}
}
