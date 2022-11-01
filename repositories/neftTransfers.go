package repositories

import (
	"shepays/db"
	"shepays/models"
)

type NeftTransferRepository struct {
	Db *db.Database
}

func (s *NeftTransferRepository) CreateNeftTransfer(w *models.NeftTransfer) error {
	return s.Db.CreateNeftTransfer_(w)
}

func (s *NeftTransferRepository) ReadNeftTransfer(userId string, NeftId string) (*models.NeftTransfer, error) {
	return s.Db.ReadNeftTransfer_(userId, NeftId)
}

func (s *NeftTransferRepository) ReadAllNeftTransfers(userId string) (*[]models.NeftTransfer, error) {
	return s.Db.ReadAllNeftTransfer_(userId)
}
