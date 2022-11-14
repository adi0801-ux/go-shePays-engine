package repositories

import (
	"shepays/db"
	"shepays/models"
)

type UserAccountCreationRepository struct {
	Db *db.Database
}

func (s *UserAccountCreationRepository) CreateUserAccount(w *models.UserAccount) error {
	return s.Db.CreateUserAccount_(w)
}

func (s *UserAccountCreationRepository) ReadUserAccount(userId string) (*models.UserAccount, error) {
	return s.Db.ReadUserAccount_(userId)
}

func (s *UserAccountCreationRepository) UpdateUserAccount(w *models.UserAccount) error {
	return s.Db.UpdateUserAccount_(w)
}

func (s *UserAccountCreationRepository) CreateOrUpdateUserAccount(w *models.UserAccount) error {
	return s.Db.CreateOrUpdateUserAccount_(w)
}
