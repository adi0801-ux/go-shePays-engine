package db

import (
	"fmt"
	"shepays/constants"
	"shepays/models"
)

func (d *Database) CreateNeftTransfer_(w *models.NeftTransfer) error {
	result := d.store.Create(&w)
	return result.Error
}

func (d *Database) ReadNeftTransfer_(userId string, neftId string) (*models.NeftTransfer, error) {
	u := &models.NeftTransfer{}
	err := d.store.Where("user_id = ? and neft_id = ?", userId, neftId).Find(u).Error
	//if err != nil {
	//	return u, err
	//}
	if u.CreatedAt.String() == constants.StartDateTime {

		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}

func (d *Database) ReadAllNeftTransfer_(userId string) (*[]models.NeftTransfer, error) {
	u := &[]models.NeftTransfer{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if err != nil {
		return u, err
	}
	if len(*u) == 0 {

		return u, fmt.Errorf(constants.NoDisputesFound)
	}
	return u, nil
}

func (d *Database) DeleteUserNeftTransfer_(userID string) error {
	userNeftTransfer := &models.NeftTransfer{}
	err := d.store.Where("user_id = ?", userID).First(userNeftTransfer).Error
	if err != nil {
		return err
	}

	err = d.store.Delete(userNeftTransfer, userID).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) CreateOrUpdateUserNeftTransfer_(
	userNeftTransfer *models.NeftTransfer) (err error) {

	if d.store.Where("user_id = ?", userNeftTransfer.UserID).Updates(&userNeftTransfer).RowsAffected == 0 {
		err = d.store.Create(&userNeftTransfer).Error
	}
	return nil
}
