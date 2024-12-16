package controller

import (
	ceknumbercontroller "main/controller/cekNumber_controller"
	"main/service"
)

type AllController struct {
	CekNum ceknumbercontroller.PhoneHandler
}

func NewAllController(service *service.AllService) AllController {
	return AllController{
		CekNum: ceknumbercontroller.NewPhoneHandler(service),
	}
}
