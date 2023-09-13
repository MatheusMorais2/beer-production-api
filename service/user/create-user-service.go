package userService

import (
	"beer-production-api/adapters/auth"
	"beer-production-api/entities/user"
	"fmt"
)

type CreateUser struct {
	UserRepository user.UserRepository
}

func NewCreateUser(userRepository user.UserRepository) *CreateUser {
	return &CreateUser{UserRepository: userRepository}
}

func (cU *CreateUser) Execute(input user.CreateUserInputDto) (*user.CreateUserOutputDto, error) {
	var err error
	userToCreate := user.NewUserApplication()
	userToCreate.Name = input.Name

	userToCreate.Email = input.Email
	userToCreate.Password, err = auth.HashPassword(input.Password)
	if err != nil {
		return &user.CreateUserOutputDto{
			ErrorMessage: "Error hashing password",
		}, err
	}

	fmt.Printf("userToCreate: %+v\n", userToCreate)
	err = userToCreate.IsValid()
	if err != nil {
		fmt.Println("entrou no erro do isvalid")
		return &user.CreateUserOutputDto{
			ErrorMessage: "Usuário inválido ou com informações faltando",
		}, err
	}

	createdUser, err := cU.UserRepository.Insert(userToCreate)
	if err != nil {
		return &user.CreateUserOutputDto{
			ErrorMessage: "Erro na criação do usuário no banco de dados",
		}, err
	}

	output := &user.CreateUserOutputDto{
		Id: createdUser.ID,
		Name: createdUser.Name,
		Email: createdUser.Email,
	}

	return output, nil
}