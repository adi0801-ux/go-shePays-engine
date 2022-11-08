package repositories

import (
	"shepays/db"
	"shepays/models"
)

type UserIntermidiateValuesRepository struct {
	Db *db.Database
}

func (s *UserIntermidiateValuesRepository) CreateUserIntermValues(w *models.UserIntermValues) error {
	return s.Db.CreateUserIntermValues_(w)
}

func (s *UserIntermidiateValuesRepository) ReadUserIntermValues(userId string) (*models.UserIntermValues, error) {
	return s.Db.ReadUserIntermValues_(userId)
}

func (s *UserIntermidiateValuesRepository) UpdateUserIntermValues(w *models.UserIntermValues) error {
	return s.Db.UpdateUserIntermValues_(w)
}

func (s *UserIntermidiateValuesRepository) CreateOrUpdateUserIntermValues(w *models.UserIntermValues) error {
	return s.Db.CreateOrUpdateUserIntermValues_(w)
}
