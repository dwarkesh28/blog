package services

import (
	userModel "blog/internal/modules/user/models"
	UserRepository "blog/internal/modules/user/repositiories"
	"blog/internal/modules/user/requests/auth"
	userResponse "blog/internal/modules/user/responses"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository UserRepository.UserRepositoryInterface
}

func New() *UserService {
	return &UserService{
		userRepository: UserRepository.New(),
	}
}

func (userService *UserService) Create(request auth.RegisterRequest) (userResponse.User, error) {
	var response userResponse.User
	var user userModel.User

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)

	if err != nil {
		return response, errors.New("error hashing password")
	}
	user.Name = request.Name
	user.Email = request.Email
	user.Password = string(hashPassword)

	newUser := userService.userRepository.Create(user)

	if newUser.ID == 0 {
		return response, errors.New("error on creating user")
	}
	return userResponse.ToUser(newUser), nil
}

func (userService *UserService) CheckUserExist(email string) bool {
	user := userService.userRepository.FindByEmail(email)
	return user.ID != 0
}

func (userService *UserService) HandleUserLogin(request auth.LoginRequest) (userResponse.User, error) {
	var response userResponse.User
	existUser := userService.userRepository.FindByEmail(request.Email)

	if existUser.ID == 0 {
		return response, errors.New("invalid credentials")
	}

	err := bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(request.Password))

	if err != nil {
		return response, errors.New("invalid credentials")
	}

	return userResponse.ToUser(existUser), nil
}
