package user

type CreateUserInputDto struct {
	Name string `json:"name"`
	Role Role `json:"role"`
	BreweryId int `json:"brewery_id"`
}