package repositories

import (
	"shepays/db"
	"shepays/models"
)

type UserCkycRepository struct {
	Db *db.Database
}

func (s *UserCkycRepository) CreateUserCkyc(w *models.UserCkyc) error {
	return s.Db.CreateUserCkyc_(w)
}

func (s *UserCkycRepository) ReadUserCkyc(userId string) (*models.UserCkyc, error) {
	return s.Db.ReadUserCkyc_(userId)
}

func (s *UserCkycRepository) ReadAllUserCkyc(userId string) (*[]models.UserCkyc, error) {
	return s.Db.ReadAllUserCkyc_(userId)
}

func (s *UserCkycRepository) CreateOrUpdateUserCkyc(w *models.UserCkyc) error {
	return s.Db.CreateOrUpdateUserCkyc_(w)
}
