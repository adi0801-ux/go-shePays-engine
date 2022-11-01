package services

import (
	"net/http"
	"shepays/models"
)

func (p *ServiceConfig) CreateUserNominees(createUsersNominees *models.UserNomineesApi) (int, interface{}, error) {
	//call api --> repository for calling happay
	for _, api := range createUsersNominees.UserNominees {
		userNominee := models.UserNominee{
			NomineeName:  api.NomineeName,
			RelationType: api.RelationType,
			UserID:       createUsersNominees.UserId,
			AddressLine1: api.Address.AddressLine1,
			City:         api.Address.City,
			Country:      api.Address.Country,
			PostalCode:   api.Address.PostalCode,
			State:        api.Address.State,
		}

		err := p.UserNomineeRep.CreateUserNominees(&userNominee)
		if err != nil {
			return http.StatusBadRequest, nil, err
		}
	}
	return http.StatusOK, nil, nil

}
