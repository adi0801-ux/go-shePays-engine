package db

import (
	"gorm.io/gorm"
	"shepays/models"
	"shepays/utils"
)

type ConnectionConfig struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
	DSN      string
}

type Database struct {
	store *gorm.DB
}

func (d *Database) RunMigrations() (err error) {
	err = d.store.AutoMigrate(&models.SavingsAccount{}, &models.UserDetail{},
		&models.UserAddress{}, &models.UserCard{}, &models.UserCkyc{},
		&models.UserNominee{}, &models.CreatePhysicalCardApiResponse{},
		&models.CreateVirtualCardApiResponse{}, &models.APILog{})
	return err
}

func (d *Database) CloseConnection() (err error) {
	conn, err := d.store.DB()

	if err != nil {
		utils.Log.Error(err)
		return
	}
	err = conn.Close()

	return

}
