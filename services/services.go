package services

import (
	"shepays/db"
	"shepays/repositories"
)

type ServiceConfig struct {
	NSDLClient              *repositories.NSDLClient
	DeviceDetailsRepo       *repositories.DeviceDetailsRepository
	CustomerDetailsRepo     *repositories.CustomerDetailsRepository
	UserKycRepo             *repositories.KycUserDocRepository
	IntermValuesRepo        *repositories.UserIntermidiateValuesRepository
	UserCardInformationRepo *repositories.UserCardCreateInformationRepository
	UserAccountRepo         *repositories.UserAccountCreationRepository
}

func CreateAllRepositoryReferences(store *db.Database) ServiceConfig {
	//create repository references
	ref := ServiceConfig{
		DeviceDetailsRepo:       &repositories.DeviceDetailsRepository{Db: store},
		CustomerDetailsRepo:     &repositories.CustomerDetailsRepository{Db: store},
		UserKycRepo:             &repositories.KycUserDocRepository{Db: store},
		IntermValuesRepo:        &repositories.UserIntermidiateValuesRepository{Db: store},
		UserCardInformationRepo: &repositories.UserCardCreateInformationRepository{Db: store},
		UserAccountRepo:         &repositories.UserAccountCreationRepository{Db: store},
	}

	return ref
}
