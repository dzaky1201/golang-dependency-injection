//go:build wireinject
// +build wireinject

package main

import (
	"golang-di/repository"
	"golang-di/service"

	"github.com/google/wire"
)

func InitializedService() *service.SimpleService {
	wire.Build(repository.NewSimpleRepository, service.NewSimpleService)
	return nil
}
