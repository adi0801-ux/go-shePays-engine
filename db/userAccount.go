package db

import (
	"fmt"
	"shepays/constants"
	"shepays/models"
)

func (d *Database) CreateUserAccount_(w *models.UserAccount) error {
	result := d.store.Create(&w)
	return result.Error
}

func (d *Database) ReadUserAccount_(userId string) (*models.UserAccount, error) {
	u := &models.UserAccount{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {

		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}

func (d *Database) UpdateUserAccount_(
	w *models.UserAccount) (err error) {

	if d.store.Where("user_id = ?", w.UserID).Updates(&w).RowsAffected == 0 {
		err = d.store.Create(&w).Error
	}
	return nil
}

func (d *Database) CreateOrUpdateUserAccount_(
	w *models.UserAccount) (err error) {

	if d.store.Where("user_id = ?", w.UserID).Updates(&w).RowsAffected == 0 {
		err = d.store.Create(&w).Error
	}
	return nil
}
