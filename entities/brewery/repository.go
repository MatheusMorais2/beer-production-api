package brewery

type CreateBreweryDBInputDTO struct {
    Name string;
    Email string;
    UserId string;
}

type BreweryRepository interface {
	Insert(brewery *CreateBreweryDBInputDTO) (*Brewery, error)
	GetBreweriesByUserId(userId string) ([]*GetUserBreweriesOutputDTO, error)
    GetUserRole(userId string, breweryId string) (string, error)
    InviteUserToBrewery(invite InviteUserInputDTO) (*Invite, error)
}