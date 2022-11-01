package repositories

import (
	"shepays/db"
	"shepays/models"
)

type CardDetailsRepository struct {
	Db *db.Database
}

func (s *CardDetailsRepository) CreateCardDetails(w *models.UserCard) error {
	return s.Db.CreateCard_(w)
}

func (s *CardDetailsRepository) ReadCard(userId string) (*models.UserCard, error) {
	return s.Db.ReadCard_(userId)
}

func (s *CardDetailsRepository) ReadAllCard(userId string) (*[]models.UserCard, error) {
	return s.Db.ReadAllCard_(userId)
}

func (s *CardDetailsRepository) SavePhysicalCardResponse(w *models.CreatePhysicalCardApiResponse) error {
	return s.Db.SavePhysicalCardDetails_(w)
}

func (s *CardDetailsRepository) UpdatePhysicalCardResponse(w *models.CreatePhysicalCardApiResponse) error {
	return s.Db.UpdatePhysicalCardDetails_(w)
}

func (s *CardDetailsRepository) SaveVirtualCardResponse(w *models.CreateVirtualCardApiResponse) error {
	return s.Db.SaveVirtualCardDetails_(w)
}

func (s *CardDetailsRepository) UpdateVirtualCardResponse(w *models.CreateVirtualCardApiResponse) error {
	return s.Db.UpdateVirtualCardDetails_(w)
}

func (s *CardDetailsRepository) ReadVirtualCardDetails(userId string) (*models.CreateVirtualCardApiResponse, error) {
	return s.Db.ReadVirtualCardDetails_(userId)
}
