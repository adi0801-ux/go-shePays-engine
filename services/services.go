package services

import (
	"shepays/db"
	"shepays/repositories"
)

type ServiceConfig struct {
	NSDLClient          *repositories.NSDLClient
	DeviceDetailsRepo   *repositories.DeviceDetailsRepository
	CustomerDetailsRepo *repositories.CustomerDetailsRepository
	UserKycRepo         *repositories.KycUserDocRepository
}

func CreateAllRepositoryReferences(store *db.Database) ServiceConfig {
	//create repository references
	ref := ServiceConfig{
		DeviceDetailsRepo:   &repositories.DeviceDetailsRepository{Db: store},
		CustomerDetailsRepo: &repositories.CustomerDetailsRepository{Db: store},
		UserKycRepo:         &repositories.KycUserDocRepository{Db: store},
	}

	return ref
}
