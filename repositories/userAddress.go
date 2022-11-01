package repositories

import (
	"shepays/db"
	"shepays/models"
)

type UserAddressRepository struct {
	Db *db.Database
}

func (s *UserAddressRepository) CreateUserAddress(w *models.UserAddress) error {
	return s.Db.CreateUserAddress_(w)
}

func (s *UserAddressRepository) ReadUserAddress(userId string) (*models.UserAddress, error) {
	return s.Db.ReadUserAddress_(userId)
}

func (s *UserAddressRepository) CreateOrUpdateUserAddress(w *models.UserAddress) error {
	return s.Db.CreateOrUpdateUserAddress_(w)
}

func (s *UserAddressRepository) GetUserAddressId(userId string) string {
	address, _ := s.Db.ReadUserAddress_(userId)
	return address.AddressID
}
