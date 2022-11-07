package repositories

import (
	"shepays/db"
	"shepays/models"
)

type DeviceDetailsRepository struct {
	Db *db.Database
}

func (s *DeviceDetailsRepository) CreateDeviceDetails(w *models.DeviceDetails) error {
	return s.Db.CreateDeviceDetails_(w)
}

func (s *DeviceDetailsRepository) ReadDeviceDetails(userId string) (*models.DeviceDetails, error) {
	return s.Db.ReadDeviceDetails_(userId)
}

func (s *DeviceDetailsRepository) UpdateDeviceDetails(w *models.DeviceDetails) error {
	return s.Db.UpdateDeviceDetails_(w)
}
