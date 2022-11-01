package repositories

import (
	"shepays/db"
	"shepays/models"
)

type UserDetailsRepository struct {
	Db *db.Database
}

func (s *UserDetailsRepository) CreateUserDetails(w *models.UserDetail) error {
	return s.Db.CreateUserDetails_(w)
}

func (s *UserDetailsRepository) ReadUserDetails(userId string) (*models.UserDetail, error) {
	return s.Db.ReadUserDetails_(userId)
}

func (s *UserDetailsRepository) UpdateUserDetails(w *models.UserDetail) error {
	return s.Db.CreateOrUpdateUserDetails_(w)
}
