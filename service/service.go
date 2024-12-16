package service

import (
	"main/repository"
	ceknumberservice "main/service/cekNumber_service"
)

type AllService struct {
	CekNum ceknumberservice.PhoneServiceInterface
}

func NewAllService(repo *repository.AllRepository) *AllService {
	return &AllService{
		CekNum: ceknumberservice.NewPhoneService(repo),
	}
}
