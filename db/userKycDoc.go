package db

import (
	"fmt"
	"shepays/constants"
	"shepays/models"
)

func (d *Database) CreateKycUserDoc_(w *models.KYCPAN) error {
	result := d.store.Create(&w)
	return result.Error
}

func (d *Database) ReadKycUserDoc_(userId string) (*models.KYCPAN, error) {
	u := &models.KYCPAN{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {

		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}

func (d *Database) UpdateKycUserDoc_(
	w *models.KYCPAN) (err error) {

	if d.store.Where("user_id = ?", w.UserID).Updates(&w).RowsAffected == 0 {
		err = d.store.Create(&w).Error
	}
	return nil
}
