package batchService

import (
	"beer-production-api/entities/batch"
	"fmt"
)

type CreateBatchStep struct {
	batchRepository batch.BatchRepository
}

func NewCreateBatchStep(batchRepository batch.BatchRepository) *CreateBatchStep {
	return &CreateBatchStep{batchRepository: batchRepository}
}

func (cbS *CreateBatchStep) Execute(input batch.CreateBatchStepInputDto) (*batch.CreateBatchStepOutputDto, error) {
	BatchStepToCreate := batch.NewBatchStepApplication()
	BatchStepToCreate.UserId = input.UserId
	BatchStepToCreate.RecipeStepId = input.RecipeStepId
	BatchStepToCreate.StartedAt = input.StartedAt
	BatchStepToCreate.BatchId = input.BatchId
	fmt.Printf("BatchStepToCreate: %+v", BatchStepToCreate)
/* 	err := BatchStepToCreate.IsValid()
	if err != nil {
		fmt.Println("entrou no erro do isvalid")
		return &batch.CreateBatchOutputDto{
			ErrorMessage: "Cervejaria inválida ou com informações faltando",
		}, err
	} */

	createdBatchStep, err := cbS.batchRepository.InsertBatchStep(BatchStepToCreate)
	if err != nil {
		return &batch.CreateBatchStepOutputDto{
			ErrorMessage: "Erro na criação da cervejaria no banco de dados",
		}, err
	}

	output := &batch.CreateBatchStepOutputDto{
		Id: createdBatchStep.ID,
		UserId: createdBatchStep.UserId,
		RecipeStepId: createdBatchStep.RecipeStepId,
		StartedAt: createdBatchStep.StartedAt,
		FinishedAt: createdBatchStep.FinishedAt,
	}

	return output, nil
}