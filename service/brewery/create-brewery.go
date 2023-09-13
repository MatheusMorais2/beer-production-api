package breweryService

import (
	"beer-production-api/entities/brewery"
	"fmt"
)

type CreateBrewery struct {
	BreweryRepository brewery.BreweryRepository
}

func NewCreateBrewery(breweryRepository brewery.BreweryRepository) *CreateBrewery {
	return &CreateBrewery{BreweryRepository: breweryRepository}
}

func (cB *CreateBrewery) Execute(input brewery.CreateBreweryInputDto) (*brewery.CreateBreweryOutputDto, error) {
	breweryToCreate := brewery.NewBreweryApplication()
	breweryToCreate.Name = input.Name
	breweryToCreate.Cnpj = input.Cnpj
	fmt.Printf("breweryToCreate: %+v\n", breweryToCreate)
	err := breweryToCreate.IsValid()
	if err != nil {
		fmt.Println("entrou no erro do isvalid")
		return &brewery.CreateBreweryOutputDto{
			ErrorMessage: "Cervejaria inválida ou com informações faltando",
		}, err
	}

	createdBrewery, err := cB.BreweryRepository.Insert(breweryToCreate)
	if err != nil {
		return &brewery.CreateBreweryOutputDto{
			ErrorMessage: "Erro na criação da cervejaria no banco de dados",
		}, err
	}

	output := &brewery.CreateBreweryOutputDto{
		Id: createdBrewery.ID,
		Name: createdBrewery.Name,
		CNPJ: createdBrewery.Cnpj,
	}

	return output, nil
}