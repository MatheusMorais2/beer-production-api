package userService

import (
	"beer-production-api/adapters/auth"
	"beer-production-api/entities/user"
	"fmt"
)

type LoginUser struct {
	UserRepository user.UserRepository
}

func NewLoginUser(userRepository user.UserRepository) *LoginUser {
	return &LoginUser{UserRepository: userRepository}
}

func (lU *LoginUser) Execute(input user.LoginUserInputDto) (*user.LoginUserOutputDto, error) {
	userToLogin := user.NewUserApplication()
	userToLogin.Password = input.Password
	userToLogin.Email = input.Email

	databaseUser, err := lU.UserRepository.GetByEmail(userToLogin.Email)
	if err != nil {
		return nil, fmt.Errorf("Error finding user")
	}

	wrongPassword := auth.CheckPasswordHash(userToLogin.Password, databaseUser.Password)
	if wrongPassword {
		return nil, fmt.Errorf("Unauthorized")
	}

	token, err := auth.GenerateUserToken(databaseUser.ID)
	if err != nil {
		return nil, fmt.Errorf("Error generating token")
	}

	output := &user.LoginUserOutputDto{
		Email: databaseUser.Email,
		Token: token,
	}

	return output, nil
}