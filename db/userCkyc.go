package db

import (
	"fmt"
	"shepays/constants"
	"shepays/models"
)

func (d *Database) CreateUserCkyc_(w *models.UserCkyc) error {
	result := d.store.Create(&w)
	return result.Error
}

func (d *Database) ReadUserCkyc_(userId string) (*models.UserCkyc, error) {
	u := &models.UserCkyc{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	//if err != nil {
	//	return u, err
	//}
	if u.CreatedAt.String() == constants.StartDateTime {

		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}

func (d *Database) ReadAllUserCkyc_(userId string) (*[]models.UserCkyc, error) {
	u := &[]models.UserCkyc{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if err != nil {
		return u, err
	}
	if len(*u) == 0 {

		return u, fmt.Errorf(constants.NoDisputesFound)
	}
	return u, nil
}

func (d *Database) DeleteUserCkyc_(userID string) error {
	userCkyc := &models.UserCkyc{}
	err := d.store.Where("user_id = ?", userID).First(userCkyc).Error
	if err != nil {
		return err
	}

	err = d.store.Delete(userCkyc, userID).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) CreateOrUpdateUserCkyc_(
	userCkyc *models.UserCkyc) (err error) {

	if d.store.Where("user_id = ?", userCkyc.UserID).Updates(&userCkyc).RowsAffected == 0 {
		err = d.store.Create(&userCkyc).Error
	}
	return nil
}
