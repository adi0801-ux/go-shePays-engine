package db

import (
	"fmt"
	"shepays/constants"
	"shepays/models"
)

func (d *Database) CreateCard_(w *models.UserCard) error {
	result := d.store.Create(&w)
	return result.Error
}

func (d *Database) ReadCard_(userId string) (*models.UserCard, error) {
	u := &models.UserCard{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	//if err != nil {
	//	return u, err
	//}
	if u.CreatedAt.String() == constants.StartDateTime {

		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}

func (d *Database) ReadAllCard_(userId string) (*[]models.UserCard, error) {
	u := &[]models.UserCard{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if err != nil {
		return u, err
	}
	if len(*u) == 0 {

		return u, fmt.Errorf(constants.NoDisputesFound)
	}
	return u, nil
}

func (d *Database) DeleteCard_(userID string) error {
	userCard := &models.UserCard{}
	err := d.store.Where("user_id = ?", userID).First(userCard).Error
	if err != nil {
		return err
	}

	err = d.store.Delete(userCard, userID).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) CreateOrUpdateCard_(
	userCard *models.UserCard) (err error) {

	if d.store.Where("user_id = ?", userCard.UserID).Updates(&userCard).RowsAffected == 0 {
		err = d.store.Create(&userCard).Error
	}
	return nil
}

func (d *Database) SavePhysicalCardDetails_(
	w *models.CreatePhysicalCardApiResponse) (err error) {

	result := d.store.Create(&w)
	return result.Error
}

func (d *Database) UpdatePhysicalCardDetails_(
	w *models.CreatePhysicalCardApiResponse) (err error) {

	if d.store.Where("user_id = ?", w.UserID).Updates(&w).RowsAffected == 0 {
		err = d.store.Create(&w).Error
	}
	return nil
}

func (d *Database) SaveVirtualCardDetails_(
	w *models.CreateVirtualCardApiResponse) (err error) {

	result := d.store.Create(&w)
	return result.Error
}

func (d *Database) UpdateVirtualCardDetails_(
	w *models.CreateVirtualCardApiResponse) (err error) {

	if d.store.Where("user_id = ?", w.UserID).Updates(&w).RowsAffected == 0 {
		err = d.store.Create(&w).Error
	}
	return nil
}

func (d *Database) ReadVirtualCardDetails_(userId string) (*models.CreateVirtualCardApiResponse, error) {
	u := &models.CreateVirtualCardApiResponse{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	//if err != nil {
	//	return u, err
	//}
	if u.CreatedAt.String() == constants.StartDateTime {

		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}
