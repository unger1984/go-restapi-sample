package users

import "awcoding.com/back/domain/uploads"

type User struct {
	Id       int             `json:"id"`
	Email    string          `json:"email"`
	AvatarId *int            `json:"avatarId"`
	Avatar   *uploads.Upload `json:"avatar"`
	Password string          `json:"-"`
}
