package ceknumberrepository

import "main/model"

type PhoneRepositoryInterface interface {
	Save(input string) bool
	GetAll() []model.PhoneNumber
	FindByNumber(number string) (*model.PhoneNumber, bool)
	Update(oldNumber, newNumber string) bool
	Delete(oldNumber string) bool
}

var data []model.PhoneNumber

type phoneRepository struct {
	DB *[]model.PhoneNumber
}

func NewPhoneRepository() PhoneRepositoryInterface {
	return &phoneRepository{
		DB: &data,
	}
}

func IsCantik(phone string) bool {
	// Deteksi triple characters (contoh: 111, 222, 333)
	for i := 0; i < len(phone)-2; i++ {
		if phone[i] == phone[i+1] && phone[i] == phone[i+2] {
			return true
		}
	}

	// Deteksi setidaknya dua double characters (contoh: 22, 33)
	doubleCount := 0
	for i := 0; i < len(phone)-1; i++ {
		if phone[i] == phone[i+1] {
			doubleCount++
			i++ // Skip character berikutnya untuk mencegah overlap
		}
	}
	return doubleCount >= 2
}

// Save adds a new phone number to the repository if it passes the IsCantik check
func (r *phoneRepository) Save(input string) bool {
	if IsCantik(input) {
		var phone model.PhoneNumber
		phone.ID = len(*r.DB) + 1
		phone.Phone = input
		*r.DB = append(*r.DB, phone)
		return true
	}
	return false
}

// GetAll returns all phone numbers in the repository
func (r *phoneRepository) GetAll() []model.PhoneNumber {
	return *r.DB
}

// FindByNumber searches for a phone number in the repository and returns it if found
func (r *phoneRepository) FindByNumber(number string) (*model.PhoneNumber, bool) {
	for _, phone := range *r.DB {
		if phone.Phone == number {
			return &phone, true
		}
	}
	return nil, false
}

// Update modifies the phone number with the given ID if it exists
func (r *phoneRepository) Update(oldNumber, newNumber string) bool {
	for i, phone := range *r.DB {
		if phone.Phone == oldNumber {
			(*r.DB)[i].Phone = newNumber
			return true
		}
	}
	return false
}

// Delete removes the phone number with the given ID from the repository
func (r *phoneRepository) Delete(oldNumber string) bool {
	for i, phone := range *r.DB {
		if phone.Phone == oldNumber {
			*r.DB = append((*r.DB)[:i], (*r.DB)[i+1:]...)
			return true
		}
	}
	return false
}
