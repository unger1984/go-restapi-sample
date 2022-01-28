package auth

import "awcoding.com/back/domain/users"

type Auth struct {
	Token string      `json:"token"`
	User  *users.User `json:"user"`
}
