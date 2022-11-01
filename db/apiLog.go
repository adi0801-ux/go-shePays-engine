package db

import "shepays/models"

func (d *Database) CreateApiLog_(w *models.APILog) error {
	result := d.store.Create(&w)
	return result.Error
}

func (d *Database) CreateOrUpdateApiLog_(
	u *models.APILog) (err error) {

	if d.store.Where("request_id = ?", u.RequestId).Updates(&u).RowsAffected == 0 {
		err = d.store.Create(&u).Error
	}
	return nil
}
