package repository

import ceknumberrepository "main/repository/cekNumber_repository"

type AllRepository struct {
	CekNum ceknumberrepository.PhoneRepositoryInterface
}

func NewAllRepo() *AllRepository {
	return &AllRepository{
		CekNum: ceknumberrepository.NewPhoneRepository(),
	}
}
