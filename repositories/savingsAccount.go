package repositories

import (
	"shepays/db"
	"shepays/models"
)

type SavingsAccountRepository struct {
	Db *db.Database
}

func (s *SavingsAccountRepository) CreateSavingsAccount(w *models.SavingsAccount) error {
	return s.Db.CreateSavingsAccount_(w)
}

func (s *SavingsAccountRepository) ReadSavingsAccount(userId string) (*models.SavingsAccount, error) {
	return s.Db.ReadSavingsAccount_(userId)
}
