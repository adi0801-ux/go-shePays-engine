package db

import (
	"fmt"
	"shepays/constants"
	"shepays/models"
)

func (d *Database) CreateUserNominees_(w *models.UserNominee) error {
	result := d.store.Create(&w)
	return result.Error
}

func (d *Database) ReadUserNominees_(userId string) (*models.UserNominee, error) {
	u := &models.UserNominee{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	//if err != nil {
	//	return u, err
	//}
	if u.CreatedAt.String() == constants.StartDateTime {

		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}

func (d *Database) ReadAllUserNominees_(userId string) (*[]models.UserNominee, error) {
	u := &[]models.UserNominee{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if err != nil {
		return u, err
	}
	if len(*u) == 0 {

		return u, fmt.Errorf(constants.NoDisputesFound)
	}
	return u, nil
}

func (d *Database) DeleteUserNominees_(userID string) error {
	userNominees := &models.UserNominee{}
	err := d.store.Where("user_id = ?", userID).First(userNominees).Error
	if err != nil {
		return err
	}

	err = d.store.Delete(userNominees, userID).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) CreateOrUpdateUserNominees_(
	UserNominee *models.UserNominee) (err error) {

	if d.store.Where("user_id = ?", UserNominee.UserID).Updates(&UserNominee).RowsAffected == 0 {
		err = d.store.Create(&UserNominee).Error
	}
	return nil
}
