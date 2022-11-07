package repositories

import (
	"shepays/db"
	"shepays/models"
)

type CustomerDetailsRepository struct {
	Db *db.Database
}

func (s *CustomerDetailsRepository) CreateCustomerDetails(w *models.CustomerDetails) error {
	return s.Db.CreateCustomerDetails_(w)
}

func (s *CustomerDetailsRepository) ReadCustomerDetails(userId string) (*models.CustomerDetails, error) {
	return s.Db.ReadCustomerDetails_(userId)
}

func (s *CustomerDetailsRepository) UpdateCustomerDetails(w *models.CustomerDetails) error {
	return s.Db.UpdateCustomerDetails_(w)
}

func (s *CustomerDetailsRepository) CreateCustomerAdditionalDetails(w *models.CustomerAdditionalInformation) error {
	return s.Db.CreateCustomerAdditionalDetails_(w)
}

func (s *CustomerDetailsRepository) ReadCustomerAdditionalDetails(userId string) (*models.CustomerAdditionalInformation, error) {
	return s.Db.ReadCustomerAdditionalDetails_(userId)
}

func (s *CustomerDetailsRepository) UpdateCustomerAdditionalDetails(w *models.CustomerAdditionalInformation) error {
	return s.Db.UpdateCustomerAdditionalDetails_(w)
}
