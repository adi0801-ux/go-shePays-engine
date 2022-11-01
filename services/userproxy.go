package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shepays/constants"
	"shepays/models"
)

func (p *ServiceConfig) CreateUserAccount(userDetail *models.UserDetail) (int, interface{}, error) {

	already, err := p.UserRep.ReadUserDetails(userDetail.UserID)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	if already.Email != "" {
		return http.StatusBadRequest, nil, fmt.Errorf("user already exists")
	}

	//call api --> repository for calling happay
	AddressId := p.UserAddressRep.GetUserAddressId(userDetail.UserID)
	if AddressId == "" {
		return http.StatusBadRequest, nil, fmt.Errorf("create user address first")
	}

	//send only required information to happay
	userDetailsApi := &models.UserDetailsApi{
		MiddleName: userDetail.MiddleName,
		Email:      userDetail.Email,
		FirstName:  userDetail.FirstName,
		Mobile:     userDetail.Mobile,
		Title:      userDetail.Title,
		Dob:        userDetail.Dob,
		Gender:     userDetail.Gender,
		LastName:   userDetail.LastName,
		AddressId:  AddressId,
	}

	response, err := p.HappayClient.SendPostRequest(constants.CreateUserEndpoint, *userDetailsApi)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	//save to db
	userDetail.HappayUserId = data["user_id"].(string)
	userDetail.KycStatus = data["kyc_status"].(string)
	userDetail.MobileVerified = data["mobile_verified"].(bool)

	err = p.UserRep.CreateUserDetails(userDetail)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return response.StatusCode, data, nil

}

func (p *ServiceConfig) ReadUserAccount(userId string) (int, interface{}, error) {
	//read from db
	data, err := p.UserRep.ReadUserDetails(userId)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, data, nil

}

func (p *ServiceConfig) UpdateUserAccount(userDetail *models.UserDetail) (int, interface{}, error) {

	//read from db
	data, err := p.UserRep.ReadUserDetails(userDetail.UserID)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	if data.Email == "" || err.Error() == constants.UserDetailNotFound {
		//	return no such user
		return http.StatusBadRequest, nil, fmt.Errorf("no such user")
	}

	response, err := p.HappayClient.SendPutRequest(constants.GenerateUpdateUserEndpoint(userDetail.UserID), *userDetail)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	if response.StatusCode != http.StatusOK {
		return http.StatusBadRequest, response, fmt.Errorf("cannot update the user")
	}

	err = p.UserRep.UpdateUserDetails(userDetail)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, userDetail, nil

}
