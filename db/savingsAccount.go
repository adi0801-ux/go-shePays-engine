package db

import (
	"fmt"
	"shepays/constants"
	"shepays/models"
)

func (d *Database) CreateSavingsAccount_(w *models.SavingsAccount) error {
	result := d.store.Create(&w)
	return result.Error
}

func (d *Database) ReadSavingsAccount_(userId string) (*models.SavingsAccount, error) {
	u := &models.SavingsAccount{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	//if err != nil {
	//	return u, err
	//}
	if u.CreatedAt.String() == constants.StartDateTime {

		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}
