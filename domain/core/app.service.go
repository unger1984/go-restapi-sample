package core

import (
	"awcoding.com/back/domain/auth"
	"awcoding.com/back/domain/users"
)

type AppServices struct {
	UserService users.UserService
	AuthService auth.AuthService
}

func NewAppService(userService users.UserService, authService auth.AuthService) *AppServices {
	return &AppServices{UserService: userService, AuthService: authService}
}
