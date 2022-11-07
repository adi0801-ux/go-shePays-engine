package db

import (
	"fmt"
	"shepays/constants"
	"shepays/models"
)

func (d *Database) CreateCustomerDetails_(w *models.CustomerDetails) error {
	result := d.store.Create(&w)
	return result.Error
}

func (d *Database) ReadCustomerDetails_(userId string) (*models.CustomerDetails, error) {
	u := &models.CustomerDetails{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {

		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}

func (d *Database) UpdateCustomerDetails_(
	w *models.CustomerDetails) (err error) {

	if d.store.Where("user_id = ?", w.UserID).Updates(&w).RowsAffected == 0 {
		err = d.store.Create(&w).Error
	}
	return nil
}

func (d *Database) CreateCustomerAdditionalDetails_(w *models.CustomerAdditionalInformation) error {
	result := d.store.Create(&w)
	return result.Error
}

func (d *Database) ReadCustomerAdditionalDetails_(userId string) (*models.CustomerAdditionalInformation, error) {
	u := &models.CustomerAdditionalInformation{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {

		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}

func (d *Database) UpdateCustomerAdditionalDetails_(
	w *models.CustomerAdditionalInformation) (err error) {

	if d.store.Where("user_id = ?", w.UserId).Updates(&w).RowsAffected == 0 {
		err = d.store.Create(&w).Error
	}
	return nil
}
