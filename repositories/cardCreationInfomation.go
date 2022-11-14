package repositories

import (
	"shepays/db"
	"shepays/models"
)

type UserCardCreateInformationRepository struct {
	Db *db.Database
}

func (s *UserCardCreateInformationRepository) CreateUserCardCreateInformation(w *models.UserCardCreateInformation) error {
	return s.Db.CreateUserCardCreateInformation_(w)
}

func (s *UserCardCreateInformationRepository) ReadUserCardCreateInformation(userId string) (*models.UserCardCreateInformation, error) {
	return s.Db.ReadUserCardCreateInformation_(userId)
}

func (s *UserCardCreateInformationRepository) UpdateUserCardCreateInformation(w *models.UserCardCreateInformation) error {
	return s.Db.UpdateUserCardCreateInformation_(w)
}

func (s *UserCardCreateInformationRepository) CreateOrUpdateUserCardCreateInformation(w *models.UserCardCreateInformation) error {
	return s.Db.CreateOrUpdateUserCardCreateInformation_(w)
}
