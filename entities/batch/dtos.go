package batch

import "time"

type CreateBatchInputDto struct {
	Name string `json:"name"`
	RecipeId string `json:"recipe_id"`
	StartDate time.Time `json:"start_date"`
}

type CreateBatchOutputDto struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name"`
	RecipeId string `json:"recipe_id"`
	StartDate time.Time `json:"start_date"`
	FinishDate time.Time `json:"finish_date,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
	StatusCode int `json:"statusCode,omitempty"`
}

type CreateBatchStepInputDto struct {
	UserId string `json:"user_id"`
	RecipeStepId string `json:"recipe_step_id"`
	BatchId string `json:"batch_id"`
	StartedAt time.Time `json:"started_at"`
}

type CreateBatchStepOutputDto struct {
	Id string `json:"id"`
	UserId string `json:"user_id"`
	RecipeStepId string `json:"recipe_step_id"`
	BatchId string `json:"batch_id"`
	StartedAt time.Time `json:"started_at"`
	FinishedAt time.Time `json:"finished_at"`
	ErrorMessage string `json:"error_message,omitempty"`
	StatusCode int `json:"statusCode,omitempty"`
}

