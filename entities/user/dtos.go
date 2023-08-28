package user

type CreateUserInputDto struct {
	Name string `json:"name"`
	Role Role `json:"role"`
	BreweryId int `json:"brewery_id"`
}

type CreateUserOutputDto struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Role Role `json:"role"`
	BreweryId int `json:"brewery_id"`
	ErrorMessage string `json:"error_message,omitempty"`
	StatusCode int `json:"statusCode,omitempty"`
}