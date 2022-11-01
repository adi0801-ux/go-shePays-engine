package db

import (
	"fmt"
	"shepays/constants"
	"shepays/models"
)

func (d *Database) CreateUserDetails_(w *models.UserDetail) error {
	result := d.store.Create(&w)
	return result.Error
}

func (d *Database) ReadUserDetails_(userId string) (*models.UserDetail, error) {
	u := &models.UserDetail{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	//if err != nil {
	//	return u, err
	//}
	if u.CreatedAt.String() == constants.StartDateTime {

		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}

func (d *Database) ReadAllUserDetails_(userId string) (*[]models.UserDetail, error) {
	u := &[]models.UserDetail{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if err != nil {
		return u, err
	}
	if len(*u) == 0 {

		return u, fmt.Errorf(constants.NoDisputesFound)
	}
	return u, nil
}

func (d *Database) DeleteUserDetails_(userID string) error {
	UserDetails := &models.UserDetail{}
	err := d.store.Where("user_id = ?", userID).First(UserDetails).Error
	if err != nil {
		return err
	}

	err = d.store.Delete(UserDetails, userID).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) CreateOrUpdateUserDetails_(
	UserDetails *models.UserDetail) (err error) {

	if d.store.Where("user_id = ?", UserDetails.UserID).Updates(&UserDetails).RowsAffected == 0 {
		err = d.store.Create(&UserDetails).Error
	}
	return nil
}
