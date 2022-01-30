package auth

import (
	"awcoding.com/back/src/domain/users"
)

type Auth struct {
	Token string      `json:"token"`
	User  *users.User `json:"user"`
}
