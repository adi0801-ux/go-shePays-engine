package repositories

import (
	"shepays/db"
	"shepays/models"
)

type DisputeRepository struct {
	Db *db.Database
}

func (s *DisputeRepository) CreateDispute(w *models.Dispute) error {
	return s.Db.CreateDispute_(w)
}

func (s *DisputeRepository) ReadDispute(userId string) (*[]models.Dispute, error) {
	return s.Db.ReadDisputes_(userId)
}

func (s *DisputeRepository) ReadAllDispute() (*[]models.Dispute, error) {
	return s.Db.ReadAllDispute_()
}
