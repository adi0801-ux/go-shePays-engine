package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shepays/constants"
	"shepays/models"
)

func (p *ServiceConfig) CreateUserAddress(userAddress *models.UserAddressInput) (int, interface{}, error) {
	//check if already not in db
	already, err := p.UserAddressRep.ReadUserAddress(userAddress.UserID)
	if already.State != "" {
		return http.StatusBadRequest, nil, fmt.Errorf("address alrady exists for the user")
	}

	//call api --> repository for calling happay

	userAddressApi := &models.UserAddressApi{
		AddressLine2: userAddress.AddressLine2,
		City:         userAddress.City,
		Country:      userAddress.Country,
		State:        userAddress.State,
		ZipCode:      userAddress.ZipCode,
		AddressLine1: userAddress.AddressLine1,
		ShippingInfo: models.ShippingInfo{
			FirstName: userAddress.FirstName,
			LastName:  userAddress.LastName,
			Mobile:    userAddress.Mobile,
		},
	}

	response, err := p.HappayClient.SendPostRequest(constants.CreateUserAddressEndpoint, *userAddressApi)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	if response.StatusCode != http.StatusCreated {
		return response.StatusCode, data, err
	}

	//save to db
	userAddressDb := models.UserAddress{
		AddressLine2: userAddress.AddressLine2,
		City:         userAddress.City,
		Country:      userAddress.Country,
		State:        userAddress.State,
		ZipCode:      userAddress.ZipCode,
		AddressLine1: userAddress.AddressLine1,
		UserID:       userAddress.UserID,
		AddressID:    data["address_id"].(string),
	}
	err = p.UserAddressRep.CreateUserAddress(&userAddressDb)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return response.StatusCode, map[string]string{"address_id": data["address_id"].(string)}, nil

}

func (p *ServiceConfig) ReadUserAddress(userId string) (int, interface{}, error) {
	//read from db
	data, err := p.UserAddressRep.ReadUserAddress(userId)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, data, nil

}

func (p *ServiceConfig) UpdateUserAddress(userAddress *models.UserAddressInput) (int, interface{}, error) {
	//read from db
	data, err := p.UserAddressRep.ReadUserAddress(userAddress.UserID)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	if data.City == "" || err.Error() == constants.UserDetailNotFound {
		//	return no such user
		return http.StatusBadRequest, nil, fmt.Errorf("no such user")
	}
	//save to db
	userAddressDb := models.UserAddress{
		AddressLine2: userAddress.AddressLine2,
		City:         userAddress.City,
		Country:      userAddress.Country,
		State:        userAddress.State,
		ZipCode:      userAddress.ZipCode,
		AddressLine1: userAddress.AddressLine1,
		UserID:       userAddress.UserID,
		AddressID:    data.AddressID,
	}
	err = p.UserAddressRep.CreateOrUpdateUserAddress(&userAddressDb)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, userAddress, nil

}
