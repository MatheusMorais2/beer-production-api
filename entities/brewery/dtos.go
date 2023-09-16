package brewery

type CreateBreweryInputDto struct {
	Name string `json:"name"`
	Cnpj string `json:"cnpj"`
	CreatorId string `json:"creator_id"`
}

type CreateBreweryOutputDto struct {
	Id string `json:"id"`
	Name string `json:"name"`
	CNPJ string `json:"cnpj"`
	ErrorMessage string `json:"error_message,omitempty"`
	StatusCode int `json:"statusCode,omitempty"`
}