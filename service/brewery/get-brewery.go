package breweryService

import (
	"beer-production-api/entities/brewery"
)

type GetBrewery struct {
	BreweryRepository brewery.BreweryRepository
}

func NewGetBrewery(breweryRepository brewery.BreweryRepository) *GetBrewery {
	return &GetBrewery{BreweryRepository: breweryRepository}
}

func (gB *GetBrewery) Execute(userId string) ([]*brewery.GetUserBreweriesOutputDTO, error) {
	userBreweries, err := gB.BreweryRepository.GetBreweriesByUserId(userId)
	if err != nil {
		return nil, err
	}

	return userBreweries, nil
}