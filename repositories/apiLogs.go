package repositories

import (
	"shepays/db"
	"shepays/models"
)

type ApiLogsRepository struct {
	Db *db.Database
}

func (s *ApiLogsRepository) CreateApiLog(w *models.APILog) error {
	return s.Db.CreateApiLog_(w)
}

func (s *ApiLogsRepository) UpdateApiLog(w *models.APILog) error {
	return s.Db.CreateOrUpdateApiLog_(w)
}
