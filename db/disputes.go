package db

import (
	"fmt"
	"shepays/constants"
	"shepays/models"
)

func (d *Database) CreateDispute_(w *models.Dispute) error {
	result := d.store.Create(&w)
	return result.Error
}

func (d *Database) ReadDisputes_(userId string) (*[]models.Dispute, error) {
	u := &[]models.Dispute{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	//if err != nil {
	//	return u, err
	//}
	if len(*u) == 0 {

		return u, fmt.Errorf(constants.NoDisputesFound)
	}
	return u, err
}

func (d *Database) ReadAllDispute_() (*[]models.Dispute, error) {
	u := &[]models.Dispute{}
	err := d.store.Find(u).Error
	if err != nil {
		return u, err
	}
	if len(*u) == 0 {

		return u, fmt.Errorf(constants.NoDisputesFound)
	}
	return u, nil
}

func (d *Database) DeleteUserDispute_(userID string) error {
	userDispute := &models.Dispute{}
	err := d.store.Where("user_id = ?", userID).First(userDispute).Error
	if err != nil {
		return err
	}

	err = d.store.Delete(userDispute, userID).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) CreateOrUpdateUserDispute_(
	userDispute *models.Dispute) (err error) {

	if d.store.Where("user_id = ?", userDispute.UserID).Updates(&userDispute).RowsAffected == 0 {
		err = d.store.Create(&userDispute).Error
	}
	return nil
}
