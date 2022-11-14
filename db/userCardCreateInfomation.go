package db

import (
	"fmt"
	"shepays/constants"
	"shepays/models"
)

func (d *Database) CreateUserCardCreateInformation_(w *models.UserCardCreateInformation) error {
	result := d.store.Create(&w)
	return result.Error
}

func (d *Database) ReadUserCardCreateInformation_(userId string) (*models.UserCardCreateInformation, error) {
	u := &models.UserCardCreateInformation{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {

		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}

func (d *Database) UpdateUserCardCreateInformation_(
	w *models.UserCardCreateInformation) (err error) {

	if d.store.Where("user_id = ?", w.UserID).Updates(&w).RowsAffected == 0 {
		err = d.store.Create(&w).Error
	}
	return nil
}

func (d *Database) CreateOrUpdateUserCardCreateInformation_(
	w *models.UserCardCreateInformation) (err error) {

	if d.store.Where("user_id = ?", w.UserID).Updates(&w).RowsAffected == 0 {
		err = d.store.Create(&w).Error
	}
	return nil
}
