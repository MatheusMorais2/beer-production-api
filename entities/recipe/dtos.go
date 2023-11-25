package recipe

type CreateRecipeInputDto struct {
	Name string `json:"name"`
	BreweryId string `json:"brewery_id"`
	Steps []Steps `json:"steps"`
	UserId string `json:"user_id"`
}

type CreateRecipeOutputDto struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	BreweryId string `json:"brewery_id,omitempty"`
	Steps []Steps `json:"steps,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
	StatusCode int `json:"statusCode,omitempty"`
}

type GetRecipesByBreweryIdInputDto struct {
	BreweryId string `param:"brewery-id"`
}

type GetRecipesByBreweryIdOutputDto struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
	StatusCode int `json:"statusCode,omitempty"`
}

type GetRecipeStepsOutputDto struct {
	Name string `json:"name,omitempty"`
	Instruction string `json:"instruction,omitempty"`
}

type GetRecipeStepsInputDto struct {
	RecipeId string `param:"recipe-id"`
	BreweryId string `param:"brewery-id"`
}