package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shepays/constants"
	"shepays/models"
)

func (p *ServiceConfig) SMSGenRequest(smsGenRequest *models.SMSGenReq) (int, interface{}, error) {
	//get device Id , uniqueId from db
	details, err := p.DeviceDetailsRepo.ReadDeviceDetails(smsGenRequest.UserId)
	if err != nil {
		return 0, nil, err
	}

	baseModel := models.SMSGenReqAPI{}

	baseModel.Channelid = p.NSDLClient.ChannelId
	baseModel.Appdtls = *p.NSDLClient.AppDtls
	baseModel.Devicedtls.Deviceid = details.DeviceId
	baseModel.Deviceidentifier = models.DeviceIdentifier{
		DeviceId:    details.DeviceId,
		DeviceUnqId: details.DeviceUniqueId,
		CustUnqId:   details.CustomerUniqueId,
	}
	baseModel.Type = smsGenRequest.Type
	baseModel.Subtype = smsGenRequest.Subtype
	baseModel.Attempt = smsGenRequest.Attempt
	baseModel.Meinumber = smsGenRequest.MeiNumber
	baseModel.Imei = smsGenRequest.Imei
	baseModel.Serialno = smsGenRequest.SerialNo
	baseModel.Simstate = smsGenRequest.SimState
	baseModel.Simoperator = smsGenRequest.SimOperator

	baseModel.Token = constants.DefaultTokenValue

	response, err := p.NSDLClient.SendPostRequest(constants.SMSGenRequestEndpoint, &baseModel)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	var data models.SMSGenReqAPIResponse
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	if data.Respcode != constants.DefaultSuccessResponseCode {
		return http.StatusBadRequest, nil, fmt.Errorf(data.Response)
	}

	statusCode, respData, err := p.VerifySmsActivation(details, smsGenRequest)
	if err != nil {
		return statusCode, respData, err
	}

	// check mobile number record added or not
	return p.CheckMobileRecord(details, smsGenRequest.UserId)

}

func (p *ServiceConfig) VerifySmsActivation(details *models.DeviceDetails, smsGenRequest *models.SMSGenReq) (int, interface{}, error) {
	//verify from nsdl
	baseModel := models.SMSActionStatusAPI{}
	baseModel.Channelid = p.NSDLClient.ChannelId
	baseModel.Appdtls = *p.NSDLClient.AppDtls
	baseModel.Devicedtls.Deviceid = details.DeviceId
	baseModel.Deviceidentifier = models.DeviceIdentifier{
		DeviceId:    details.DeviceId,
		DeviceUnqId: details.DeviceUniqueId,
		CustUnqId:   details.CustomerUniqueId,
	}
	baseModel.Type = smsGenRequest.Type
	baseModel.Subtype = smsGenRequest.Subtype

	baseModel.Token = constants.DefaultTokenValue

	response, err := p.NSDLClient.SendPostRequest(constants.SMSStatusActionEndpoint, &baseModel)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	var data models.SMSActionStatusAPIResponse
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	if data.Respcode != constants.DefaultSuccessResponseCode {
		return http.StatusBadRequest, nil, fmt.Errorf(data.Response)
	}

	return response.StatusCode, data, nil
}

func (p *ServiceConfig) CheckMobileRecord(details *models.DeviceDetails, UserId string) (int, interface{}, error) {
	customerDetails, err := p.CustomerDetailsRepo.ReadCustomerDetails(UserId)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	baseModel := models.CustomerRecordMobileNumberCheck{}
	baseModel.Channelid = p.NSDLClient.ChannelId
	baseModel.Appdtls = *p.NSDLClient.AppDtls
	baseModel.Devicedtls.Deviceid = details.DeviceId
	baseModel.Deviceidentifier = models.DeviceIdentifier{
		DeviceId:    details.DeviceId,
		DeviceUnqId: details.DeviceUniqueId,
		CustUnqId:   details.CustomerUniqueId,
	}
	baseModel.Searchtype = "MOBILENO"
	baseModel.Searchstr = customerDetails.MobileCountryCode + customerDetails.Msisdn

	baseModel.Token = constants.DefaultTokenValue

	response, err := p.NSDLClient.SendPostRequest(constants.CheckCustomerMobileNumber, &baseModel)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	var data models.CustomerRecordMobileNumberCheckResponse
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	if data.Respcode != constants.DefaultSuccessResponseCode {
		return http.StatusBadRequest, nil, fmt.Errorf(data.Response)
	}

	return response.StatusCode, nil, nil

}
