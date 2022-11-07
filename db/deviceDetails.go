package db

import (
	"fmt"
	"shepays/constants"
	"shepays/models"
)

func (d *Database) CreateDeviceDetails_(w *models.DeviceDetails) error {
	result := d.store.Create(&w)
	return result.Error
}

func (d *Database) ReadDeviceDetails_(userId string) (*models.DeviceDetails, error) {
	u := &models.DeviceDetails{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {

		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}

func (d *Database) UpdateDeviceDetails_(
	w *models.DeviceDetails) (err error) {

	if d.store.Where("user_id = ?", w.UserID).Updates(&w).RowsAffected == 0 {
		err = d.store.Create(&w).Error
	}
	return nil
}
