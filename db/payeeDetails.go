package db

import (
	"fmt"
	"shepays/constants"
	"shepays/models"
)

func (d *Database) CreatePayeeDetail_(w *models.PayeeDetail) error {
	result := d.store.Create(&w)
	return result.Error
}

func (d *Database) ReadAllUserPayeeDetail_(userId string) (*[]models.PayeeDetail, error) {
	u := &[]models.PayeeDetail{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if err != nil {
		return u, err
	}
	if len(*u) == 0 {

		return u, fmt.Errorf(constants.NoDisputesFound)
	}
	return u, nil
}

func (d *Database) DeleteUserPayeeDetail_(userID string) error {
	userPayeeDetail := &models.PayeeDetail{}
	err := d.store.Where("user_id = ?", userID).First(userPayeeDetail).Error
	if err != nil {
		return err
	}

	err = d.store.Delete(userPayeeDetail, userID).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) CreateOrUpdateUserPayeeDetail_(
	userPayeeDetail *models.PayeeDetail) (err error) {

	if d.store.Where("user_id = ?", userPayeeDetail.UserID).Updates(&userPayeeDetail).RowsAffected == 0 {
		err = d.store.Create(&userPayeeDetail).Error
	}
	return nil
}
