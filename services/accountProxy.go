package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shepays/constants"
	"shepays/models"
	"shepays/utils"
)

func (p *ServiceConfig) CreateSavingsAccount(userID string) (int, interface{}, error) {
	//validate all requirements for the user ID
	//check if account already does not exists
	alreadyAccount, _ := p.AccountRep.ReadSavingsAccount(userID)
	if alreadyAccount.BankAccountNumber == "" {
		return http.StatusBadRequest, nil, fmt.Errorf(constants.AccountAlreadyCreated)
	}

	//get doc details
	ckyc, err := p.CkycRep.ReadUserCkyc(userID)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	nomineesDb, err := p.UserNomineeRep.ReadAllUserNominees(userID)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	users, err := p.UserRep.ReadUserDetails(userID)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	//make model to create savings account
	TotalValue := 100
	TotalLength := len(*nomineesDb)
	var nominees []models.NomineesSavingsAccount
	for index, nominee := range *nomineesDb {
		var val int
		if index == TotalLength-1 {
			val = TotalValue - index
		} else {
			val = 1
		}
		newNominee := models.NomineesSavingsAccount{
			Address: models.AddressSavingsAccount{
				AddressLine1: nominee.AddressLine1,
				City:         nominee.City,
				Country:      nominee.Country,
				PostalCode:   nominee.PostalCode,
				State:        nominee.State,
			},
			NomineeName:  nominee.NomineeName,
			RelationType: nominee.RelationType,
			Value:        val,
		}
		nominees = append(nominees, newNominee)
	}

	savingsAccount := &models.CreateSavingsAccount{
		DocDetails: models.DocDetails{
			DocNumber:    ckyc.DocID,
			IssueDate:    ckyc.IssueDate,
			PlaceOfIssue: ckyc.PlaceOfIssue,
		},
		NomineeDetails: nominees,
		RelatedDetails: models.RelatedDetails{
			EmploymentStatus: users.EmploymentStatus,
			MaritalStatus:    users.MaritalStatus,
			Nationality:      users.Nationality,
		},
		UserId:      users.HappayUserId,
		AccountType: constants.DefaultAccountType,
	}

	response, err := p.HappayClient.SendPostRequest(constants.CreateSavingsAccountEndpoint, *savingsAccount)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	if response.StatusCode != http.StatusCreated {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	var data models.SavingsAccountApiResponse
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	//save to db
	account := models.SavingsAccount{
		UserID:            userID,
		AccountId:         data.AccountId,
		Ifsc:              data.Ifsc,
		BankAccountNumber: data.BankAccountNumber,
	}

	err = p.AccountRep.CreateSavingsAccount(&account)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	data.HappayUserId = data.UserId
	data.UserId = userID

	return response.StatusCode, data, nil
}

func (p *ServiceConfig) InitiateNeftTransfer(initiateNeftTransfer *models.InitiateNeftTransfer) (int, interface{}, error) {
	//call api --> repository for calling happay

	response, err := p.HappayClient.SendPostRequest(constants.InitiateNEFTTransferEndpoint, *initiateNeftTransfer)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return response.StatusCode, data, nil

}

func (p *ServiceConfig) ValidateNeftTransfer(validateNeftTransfer *models.ValidateNeftTransafer) (int, interface{}, error) {
	//call api --> repository for calling happay

	response, err := p.HappayClient.SendPostRequest(constants.ValidateNEFTTransferEndpoint, *validateNeftTransfer)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return response.StatusCode, data, nil

}

func (p *ServiceConfig) CheckPaymentStatus(paymentId string) (int, interface{}, error) {
	//call api --> repository for calling happay

	response, err := p.HappayClient.SendGetRequest(constants.CheckTransferStatusEndpoint+paymentId, nil)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return response.StatusCode, data, nil
}
