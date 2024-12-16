package ceknumberservice

import (
	"main/model"
	"main/repository"
)

type PhoneServiceInterface interface {
	FilterCantikNumbers(input string) string
	GetAllNumbers() []model.PhoneNumber
	FindNumber(input string) (*model.PhoneNumber, string)
	UpdateNumber(oldNumber, newNumber string) string
	DeleteNumber(oldNumber string) string
}

type phoneService struct {
	Repo *repository.AllRepository
}

func NewPhoneService(repo *repository.AllRepository) PhoneServiceInterface {
	return &phoneService{
		Repo: repo,
	}
}

// FilterCantikNumbers checks if the number is a "cantik" number and saves it if true
func (s *phoneService) FilterCantikNumbers(input string) string {
	if s.Repo.CekNum.Save(input) {
		return "Nomor Cantik dan berhasil disimpan"
	}
	return "Nomor tidak cantik dan tidak disimpan"
}

// GetAllNumbers retrieves all phone numbers from the repository
func (s *phoneService) GetAllNumbers() []model.PhoneNumber {
	return s.Repo.CekNum.GetAll()
}

// FindNumber searches for a phone number and returns its details if found
func (s *phoneService) FindNumber(input string) (*model.PhoneNumber, string) {
	if phone, found := s.Repo.CekNum.FindByNumber(input); found {
		return phone, "Nomor ditemukan"
	}
	return nil, "Nomor tidak ditemukan"
}

// UpdateNumber updates a phone number by its ID
func (s *phoneService) UpdateNumber(oldNumber, newNumber string) string {
	if s.Repo.CekNum.Update(oldNumber, newNumber) {
		return "Nomor berhasil diperbarui"
	}
	return "Nomor gagal diperbarui Karena Nomor tidak terdaftar"
}

// DeleteNumber deletes a phone number by its ID
func (s *phoneService) DeleteNumber(oldNumber string) string {
	if s.Repo.CekNum.Delete(oldNumber) {
		return "Nomor berhasil dihapus"
	}
	return "Nomor gagal dihapus Karena nomor tidak terdaftar"
}
