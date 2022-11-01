package db

import (
	"fmt"
	"shepays/constants"
	"shepays/models"
)

func (d *Database) CreateUserAddress_(w *models.UserAddress) error {
	result := d.store.Create(&w)
	return result.Error
}

func (d *Database) ReadUserAddress_(userId string) (*models.UserAddress, error) {
	u := &models.UserAddress{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	//if err != nil {
	//	return u, err
	//}
	if u.CreatedAt.String() == constants.StartDateTime {

		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}

func (d *Database) ReadAllUserAddress_(userId string) (*[]models.UserAddress, error) {
	u := &[]models.UserAddress{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if err != nil {
		return u, err
	}
	if len(*u) == 0 {

		return u, fmt.Errorf(constants.NoDisputesFound)
	}
	return u, nil
}

func (d *Database) DeleteUserAddress_(userID string) error {
	userAddress := &models.UserAddress{}
	err := d.store.Where("user_id = ?", userID).First(userAddress).Error
	if err != nil {
		return err
	}

	err = d.store.Delete(userAddress, userID).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) CreateOrUpdateUserAddress_(
	userAddress *models.UserAddress) (err error) {

	if d.store.Where("user_id = ?", userAddress.UserID).Updates(&userAddress).RowsAffected == 0 {
		err = d.store.Create(&userAddress).Error
	}
	return nil
}
