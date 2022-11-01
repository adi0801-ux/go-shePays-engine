package services

import (
	"shepays/db"
	"shepays/repositories"
)

type ServiceConfig struct {
	CardRep        *repositories.CardDetailsRepository
	AccountRep     *repositories.SavingsAccountRepository
	HappayClient   *repositories.HappyClient
	CkycRep        *repositories.UserCkycRepository
	UserRep        *repositories.UserDetailsRepository
	UserAddressRep *repositories.UserAddressRepository
	UserNomineeRep *repositories.UserNomineesRepository
}

func CreateAllRepositoryReferences(store *db.Database) ServiceConfig {
	//create repository references
	ref := ServiceConfig{
		CardRep:        &repositories.CardDetailsRepository{Db: store},
		AccountRep:     &repositories.SavingsAccountRepository{Db: store},
		CkycRep:        &repositories.UserCkycRepository{Db: store},
		UserRep:        &repositories.UserDetailsRepository{Db: store},
		UserAddressRep: &repositories.UserAddressRepository{Db: store},
		UserNomineeRep: &repositories.UserNomineesRepository{Db: store},
	}

	return ref
}
