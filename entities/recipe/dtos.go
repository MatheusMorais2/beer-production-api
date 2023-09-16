package recipe

type CreateRecipeInputDto struct {
	Name string `json:"name"`
	BreweryId string `json:"brewery_id"`
	Steps []Steps `json:"steps"`
}

type CreateRecipeOutputDto struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	BreweryId string `json:"brewery_id,omitempty"`
	Steps []Steps `json:"steps,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
	StatusCode int `json:"statusCode,omitempty"`
}