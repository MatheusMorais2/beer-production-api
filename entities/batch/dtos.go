package batch

import "time"

type CreateBatchInputDto struct {
	Name string `json:"name"`
	RecipeId int `json:"recipe_id"`
	StartDate time.Time `json:"start_date"`
}

type CreateBatchOutputDto struct {
	Id int `json:"id,omitempty"`
	Name string `json:"name"`
	RecipeId int `json:"recipe_id"`
	StartDate time.Time `json:"start_date"`
	FinishDate time.Time `json:"finish_date,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
	StatusCode int `json:"statusCode,omitempty"`
}

type CreateBatchStepInputDto struct {
	UserId int `json:"user_id"`
	RecipeStepId int `json:"recipe_step_id"`
	BatchId int `json:"batch_id"`
	StartedAt time.Time `json:"started_at"`
}

type CreateBatchStepOutputDto struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	RecipeStepId int `json:"recipe_step_id"`
	BatchId int `json:"batch_id"`
	StartedAt time.Time `json:"started_at"`
	FinishedAt time.Time `json:"finished_at"`
	ErrorMessage string `json:"error_message,omitempty"`
	StatusCode int `json:"statusCode,omitempty"`
}

