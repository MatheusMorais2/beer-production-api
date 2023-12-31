package batchService

import (
	"beer-production-api/entities/batch"
)

type CreateBatch struct {
	batchRepository batch.BatchRepository
}

func NewCreateBatch(batchRepository batch.BatchRepository) *CreateBatch {
	return &CreateBatch{batchRepository: batchRepository}
}

func (cB *CreateBatch) Execute(input batch.CreateBatchInputDto) (*batch.CreateBatchOutputDto, error) {
	BatchToCreate := batch.NewBatchApplication()
	BatchToCreate.Name = input.Name
	BatchToCreate.RecipeId = input.RecipeId
	BatchToCreate.StartDate = input.StartDate
	err := BatchToCreate.IsValid()
	if err != nil {
		return &batch.CreateBatchOutputDto{
			ErrorMessage: "Cervejaria inválida ou com informações faltando",
		}, err
	}

	createdBatch, err := cB.batchRepository.Insert(BatchToCreate)
	if err != nil {
		return &batch.CreateBatchOutputDto{
			ErrorMessage: "Erro na criação da cervejaria no banco de dados",
		}, err
	}

	output := &batch.CreateBatchOutputDto{
		Id: createdBatch.ID,
		Name: createdBatch.Name,
		RecipeId: createdBatch.ID,
		StartDate: createdBatch.StartDate,
	}

	return output, nil
}