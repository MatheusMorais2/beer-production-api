package batchService

import (
	"beer-production-api/entities/batch"
	"fmt"
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
	fmt.Printf("BatchToCreate: %+v\n", BatchToCreate)
	err := BatchToCreate.IsValid()
	if err != nil {
		fmt.Println("entrou no erro do isvalid")
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