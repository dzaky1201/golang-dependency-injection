package service

import (
	"fmt"
	"golang-di/repository"
)

type SimpleService struct {
	SimpleRepository repository.SimpleRepository
}

func NewSimpleService(repo *repository.SimpleRepository) *SimpleService {
	return &SimpleService{
		SimpleRepository: *repo,
	}
}

func (service *SimpleService)GetName()  {
	fmt.Println("ini contoh service sudah bisa")
}