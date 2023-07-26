package services

import (
	"blog/internal/modules/user/requests/auth"
	userResponse "blog/internal/modules/user/responses"
)

type UserServiceInterface interface {
	Create(request auth.RegisterRequest) (userResponse.User, error)
	CheckUserExist(email string) bool
	HandleUserLogin(request auth.LoginRequest) (userResponse.User, error)
}
