package batch

type BatchRepository interface {
	Insert(batch *Batch) (*Batch, error)
	InsertBatchStep(batchStepToInsert *BatchStep) (*BatchStep, error)
}