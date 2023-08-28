package brewery

type CreateBreweryInputDto struct {
	Name string `json:"name"`
	Cnpj string `json:"cnpj"`
}

type CreateBreweryOutputDto struct {
	Id int `json:"id"`
	Name string `json:"name"`
	CNPJ string `json:"cnpj"`
	ErrorMessage string `json:"error_message,omitempty"`
	StatusCode int `json:"statusCode,omitempty"`
}