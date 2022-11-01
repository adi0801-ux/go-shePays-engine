package repositories

import (
	"shepays/db"
	"shepays/models"
)

type UserNomineesRepository struct {
	Db *db.Database
}

func (s *UserNomineesRepository) CreateUserNominees(w *models.UserNominee) error {
	return s.Db.CreateUserNominees_(w)
}

func (s *UserNomineesRepository) ReadUserNominees(userId string) (*models.UserNominee, error) {
	return s.Db.ReadUserNominees_(userId)
}

func (s *UserNomineesRepository) ReadAllUserNominees(userId string) (*[]models.UserNominee, error) {
	return s.Db.ReadAllUserNominees_(userId)
}

func (s *UserNomineesRepository) UpdateUserNominees(w *models.UserNominee) error {
	return s.Db.CreateOrUpdateUserNominees_(w)
}
