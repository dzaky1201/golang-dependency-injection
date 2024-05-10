package repository

type SimpleRepository struct{}

func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{}
}