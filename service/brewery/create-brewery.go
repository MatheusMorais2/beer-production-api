package breweryService

import (
	"beer-production-api/entities/brewery"
	"beer-production-api/entities/user"
)

type CreateBrewery struct {
	BreweryRepository brewery.BreweryRepository
	UserRepository user.UserRepository
}

func NewCreateBrewery(breweryRepository brewery.BreweryRepository, userRepository user.UserRepository) *CreateBrewery {
	return &CreateBrewery{BreweryRepository: breweryRepository, UserRepository: userRepository}
}

func (cB *CreateBrewery) Execute(input brewery.CreateBreweryInputDto) (*brewery.CreateBreweryOutputDto, error) {
	breweryToCreate := brewery.NewBreweryApplication()
	breweryToCreate.Name = input.Name
	breweryToCreate.Email = input.Email
	
	err := breweryToCreate.IsValid()
	if err != nil {
		return &brewery.CreateBreweryOutputDto{
			ErrorMessage: "Cervejaria inválida ou com informações faltando",
			}, err
		}
		
	user, err := cB.UserRepository.GetById(input.UserId)
	if err != nil {
		return &brewery.CreateBreweryOutputDto{
			ErrorMessage: "Usuário não encontrado",
			}, err
		}

	inputBrewery := brewery.CreateBreweryDBInputDTO{
		Name: breweryToCreate.Name,
		Email: breweryToCreate.Email,
		UserId: user.ID,
	}

	createdBrewery, err := cB.BreweryRepository.Insert(&inputBrewery)
	if err != nil {
		return &brewery.CreateBreweryOutputDto{
			ErrorMessage: "Erro na criação da cervejaria no banco de dados",
		}, err
	}

	output := &brewery.CreateBreweryOutputDto{
		Id: createdBrewery.ID,
		Name: createdBrewery.Name,
		Email: createdBrewery.Email,
	}

	return output, nil
}