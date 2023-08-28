package userService

import (
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
	userToCreate := user.NewUserApplication()
	userToCreate.Name = input.Name
	userToCreate.Role = input.Role
	userToCreate.BreweryId = input.BreweryId
	fmt.Printf("userToCreate: %+v", userToCreate)
	err := userToCreate.IsValid()
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
		Role: createdUser.Role,
		BreweryId: createdUser.BreweryId,
	}

	return output, nil
}