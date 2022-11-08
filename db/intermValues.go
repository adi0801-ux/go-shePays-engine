package db

import (
	"fmt"
	"shepays/constants"
	"shepays/models"
)

func (d *Database) CreateUserIntermValues_(w *models.UserIntermValues) error {
	result := d.store.Create(&w)
	return result.Error
}

func (d *Database) ReadUserIntermValues_(userId string) (*models.UserIntermValues, error) {
	u := &models.UserIntermValues{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {

		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}

func (d *Database) UpdateUserIntermValues_(
	w *models.UserIntermValues) (err error) {

	if d.store.Where("user_id = ?", w.UserID).Updates(&w).RowsAffected == 0 {
		err = d.store.Create(&w).Error
	}
	return nil
}

func (d *Database) CreateOrUpdateUserIntermValues_(
	w *models.UserIntermValues) (err error) {

	if d.store.Where("user_id = ?", w.UserID).Updates(&w).RowsAffected == 0 {
		err = d.store.Create(&w).Error
	}
	return nil
}
