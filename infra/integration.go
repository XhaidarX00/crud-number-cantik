package infra

import (
	"main/controller"
	"main/repository"
	"main/service"
)

type IntegrationContext struct {
	Ctl controller.AllController
}

func NewIntegrateContext() (*IntegrationContext, error) {
	repo := repository.NewAllRepo()

	service := service.NewAllService(repo)

	handler := controller.NewAllController(service)

	return &IntegrationContext{
		Ctl: handler,
	}, nil
}
