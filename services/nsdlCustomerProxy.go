package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shepays/constants"
	"shepays/models"
	"shepays/utils"
)

func (p *ServiceConfig) CustomerInformation(customerInfo *models.CustomerInformation) (int, interface{}, error) {
	//get device Id , uniqueId from db
	details, err := p.DeviceDetailsRepo.ReadDeviceDetails(customerInfo.UserId)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	baseModel := models.CustomerDeviceApi{}

	baseModel.Mbappdetails.Channelid = p.NSDLClient.ChannelId
	baseModel.Mbappdetails.Appdtls = *p.NSDLClient.AppDtls
	baseModel.Mbappdetails.Devicedtls.Deviceid = details.DeviceId
	baseModel.Mbappdetails.Deviceidentifier = models.DeviceIdentifier{
		DeviceId:    details.DeviceId,
		DeviceUnqId: details.DeviceUniqueId,
	}
	baseModel.Mbappdetails.Custprofiledtls.CustName = customerInfo.CustomerName
	baseModel.Mbappdetails.Custprofiledtls.CustNameF = customerInfo.CustomerFName
	baseModel.Mbappdetails.Custprofiledtls.CustNameM = customerInfo.CustomerMName
	baseModel.Mbappdetails.Custprofiledtls.CustNameL = customerInfo.CustomerLName
	baseModel.Mbappdetails.Custprofiledtls.CustNickname = customerInfo.CustomerNickname
	baseModel.Mbappdetails.Custprofiledtls.Emailid = customerInfo.EmailId
	baseModel.Mbappdetails.Custprofiledtls.Msisdn = customerInfo.Msisdn
	baseModel.Mbappdetails.Custprofiledtls.Mobilecountrycode = customerInfo.MobileCountryCode
	baseModel.Mbappdetails.Custprofiledtls.Gender = customerInfo.Gender

	baseModel.Token = constants.DefaultTokenValue

	response, err := p.NSDLClient.SendPostRequest(constants.CustomerDeviceRegisterEndpoint, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	var data models.CustomerDeviceApiResponse
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	if data.Respcode != constants.DefaultSuccessResponseCode {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, fmt.Errorf(data.Response)
	}

	//save to db
	details.CustomerUniqueId = data.Deviceidentifier.Custunqid
	err = p.DeviceDetailsRepo.UpdateDeviceDetails(details)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	//save customer details

	customerDetails := models.CustomerDetails{
		CustomerNickname:  customerInfo.CustomerNickname,
		CustomerName:      customerInfo.CustomerName,
		CustomerFName:     customerInfo.CustomerFName,
		CustomerMName:     customerInfo.CustomerMName,
		CustomerLName:     customerInfo.CustomerLName,
		EmailId:           customerInfo.EmailId,
		Msisdn:            customerInfo.Msisdn,
		MobileCountryCode: customerInfo.MobileCountryCode,
		Gender:            customerInfo.Gender,
		UserID:            customerInfo.UserId,
	}

	err = p.CustomerDetailsRepo.CreateCustomerDetails(&customerDetails)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	return response.StatusCode, map[string]string{"device_unique_id": details.DeviceUniqueId, "device_id": details.DeviceId, "customer_unique_id": details.CustomerUniqueId}, nil

}

func (p *ServiceConfig) CustomerAdditionalInformation(customerAdditionalInformation *models.CustomerAdditionalInformation) (int, interface{}, error) {

	//get device Id , uniqueId from db
	details, err := p.DeviceDetailsRepo.ReadDeviceDetails(customerAdditionalInformation.UserId)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	//check pincode
	statusCode, data, err := p.ValidatePincode(customerAdditionalInformation, details)
	if err != nil {
		return statusCode, data, err
	}

	//call api to set information
	statusCode, data, err = p.SetCustomerAddditionalInformation(customerAdditionalInformation, details)
	if err != nil {
		return statusCode, data, err
	}

	//	save to db
	err = p.CustomerDetailsRepo.CreateCustomerAdditionalDetails(customerAdditionalInformation)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	//call consent api for consent asking
	statusCode, data, err = p.CustomerConsentAsking(details)
	if err != nil {
		return statusCode, data, err
	}

	return statusCode, nil, err

}

func (p *ServiceConfig) ValidatePincode(customerAdditionalInformation *models.CustomerAdditionalInformation, details *models.DeviceDetails) (int, interface{}, error) {

	baseModel := models.ValidatePincode{}
	baseModel.Channelid = p.NSDLClient.ChannelId
	baseModel.Appdtls = *p.NSDLClient.AppDtls
	baseModel.Devicedtls.Deviceid = details.DeviceId
	baseModel.Deviceidentifier = models.DeviceIdentifier{
		DeviceId:    details.DeviceId,
		DeviceUnqId: details.DeviceUniqueId,
		CustUnqId:   details.CustomerUniqueId,
	}
	baseModel.Pincode = customerAdditionalInformation.PinCode

	baseModel.Token = constants.DefaultTokenValue

	response, err := p.NSDLClient.SendPostRequest(constants.ValidatePincodeEndpoint, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	var data models.ValidatePincodeResponse
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	if data.Respcode != constants.DefaultSuccessResponseCode {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, fmt.Errorf(data.Response)
	}
	if data.RespMessage != constants.DefaultPincodeValidated {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, fmt.Errorf(data.RespMessage)
	}

	return response.StatusCode, nil, nil
}

func (p *ServiceConfig) SetCustomerAddditionalInformation(customerAdditionalInformation *models.CustomerAdditionalInformation, details *models.DeviceDetails) (int, interface{}, error) {

	customerDetails, err := p.CustomerDetailsRepo.ReadCustomerDetails(customerAdditionalInformation.UserId)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	baseModel := models.CustomerAdditionalInformationAPI{}
	baseModel.Channelid = p.NSDLClient.ChannelId
	baseModel.Appdtls = *p.NSDLClient.AppDtls
	baseModel.Devicedtls.Deviceid = details.DeviceId
	baseModel.Deviceidentifier = models.DeviceIdentifier{
		DeviceId:    details.DeviceId,
		DeviceUnqId: details.DeviceUniqueId,
		CustUnqId:   details.CustomerUniqueId,
	}

	baseModel.Custmindtls.Customername = customerDetails.CustomerName
	baseModel.Custmindtls.Custemail = customerDetails.EmailId
	baseModel.Custmindtls.Dateofbirth = customerAdditionalInformation.DOB
	baseModel.Custmindtls.Gender = customerDetails.Gender
	baseModel.Custmindtls.Salutation = customerAdditionalInformation.Salutation
	baseModel.Custmindtls.Country = customerAdditionalInformation.Country
	baseModel.Custmindtls.Pincode = customerAdditionalInformation.PinCode
	baseModel.Custmindtls.State = customerAdditionalInformation.State
	baseModel.Custmindtls.City = customerAdditionalInformation.City
	baseModel.Custmindtls.Msisdn = customerDetails.Msisdn
	baseModel.Custmindtls.Mobcntrycode = customerDetails.MobileCountryCode

	baseModel.Token = constants.DefaultTokenValue

	response, err := p.NSDLClient.SendPostRequest(constants.SetAdditionalInfomrmationEndpoint, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	var data models.CustomerAdditionalInformationAPIResponse
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	if data.Respcode != constants.DefaultSuccessResponseCode {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, fmt.Errorf(data.Response)
	}

	return response.StatusCode, data, nil

}

func (p *ServiceConfig) CustomerConsentAsking(details *models.DeviceDetails) (int, interface{}, error) {

	baseModel := models.ConsentAskingAPI{}

	baseModel.Channelid = p.NSDLClient.ChannelId
	baseModel.Appdtls = *p.NSDLClient.AppDtls
	baseModel.Devicedtls.Deviceid = details.DeviceId
	baseModel.Deviceidentifier = models.DeviceIdentifier{
		DeviceId:    details.DeviceId,
		DeviceUnqId: details.DeviceUniqueId,
		CustUnqId:   details.CustomerUniqueId,
	}
	baseModel.Consentflg = "Y"

	baseModel.Token = constants.DefaultTokenValue

	response, err := p.NSDLClient.SendPostRequest(constants.ConsentAskingEndpoint, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	var data models.ConsentAskingAPIResponse
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	if data.Respcode != constants.DefaultSuccessResponseCode {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, fmt.Errorf(data.Response)
	}

	return response.StatusCode, data, nil

}

func (p *ServiceConfig) SetMPINProxy(setMPIN *models.SetMPIN) (int, interface{}, error) {

	//get device Id , uniqueId from db
	details, err := p.DeviceDetailsRepo.ReadDeviceDetails(setMPIN.UserId)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	baseModel := models.SetMPINAPI{}

	baseModel.Channelid = p.NSDLClient.ChannelId
	baseModel.Appdtls = *p.NSDLClient.AppDtls
	baseModel.Devicedtls.Deviceid = details.DeviceId
	baseModel.Deviceidentifier = models.DeviceIdentifier{
		DeviceId:    details.DeviceId,
		DeviceUnqId: details.DeviceUniqueId,
		CustUnqId:   details.CustomerUniqueId,
	}
	baseModel.Creddtls.Credtype = "MPIN"
	baseModel.Creddtls.Credval = setMPIN.CredVal
	baseModel.Creddtls.Credcategory = "LOGIN"
	baseModel.Creddtls.Credsaltkey = setMPIN.CredSaltKey

	baseModel.Token = constants.DefaultTokenValue

	response, err := p.NSDLClient.SendPostRequest(constants.SetMPINEndpoint, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	var data interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	if data.(map[string]interface{})["respcode"].(string) != constants.DefaultSuccessResponseCode {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, fmt.Errorf(data.(map[string]interface{})["response"].(string))
	}

	return response.StatusCode, data, nil

}

func (p *ServiceConfig) GetOTPProxy(user *models.UserId) (int, interface{}, error) {

	//get device Id , uniqueId from db
	details, err := p.DeviceDetailsRepo.ReadDeviceDetails(user.UserId)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	userDetails, err := p.CustomerDetailsRepo.ReadCustomerDetails(user.UserId)

	baseModel := models.SMSGenrationAPI{}

	baseModel.Channelid = p.NSDLClient.ChannelId
	baseModel.Appdtls = *p.NSDLClient.AppDtls
	baseModel.Devicedtls.Deviceid = details.DeviceId
	baseModel.Deviceidentifier = models.DeviceIdentifier{
		DeviceId:    details.DeviceId,
		DeviceUnqId: details.DeviceUniqueId,
		CustUnqId:   details.CustomerUniqueId,
	}

	baseModel.Mobileno = userDetails.Msisdn
	baseModel.Emailid = userDetails.EmailId
	baseModel.Sentmode = "SMS"
	baseModel.Type = "ESIGNOTP"
	baseModel.Subtype = "OTPAUTHENTICATION"
	baseModel.Attempt = "1"

	baseModel.Token = constants.DefaultTokenValue

	response, err := p.NSDLClient.SendPostRequest(constants.SMSGenrationEndpoint, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	var data models.SMSGenrationAPIResponse

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	if data.Respcode != constants.DefaultSuccessResponseCode {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, fmt.Errorf(data.Response)
	}

	//	save id
	values, _ := p.IntermValuesRepo.ReadUserIntermValues(user.UserId)
	values.OtpRequiredId = data.OtpDtls.Otpreqid
	values.UserID = user.UserId

	err = p.IntermValuesRepo.CreateUserIntermValues(values)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	return response.StatusCode, data, nil

}

func (p *ServiceConfig) VerifyOTPProxy(otp *models.OTPVerify) (int, interface{}, error) {

	//get device Id , uniqueId from db
	details, err := p.DeviceDetailsRepo.ReadDeviceDetails(otp.UserId)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	userDetails, err := p.CustomerDetailsRepo.ReadCustomerDetails(otp.UserId)

	values, _ := p.IntermValuesRepo.ReadUserIntermValues(otp.UserId)

	baseModel := models.SMSVerifyAPI{}

	baseModel.Channelid = p.NSDLClient.ChannelId
	baseModel.Appdtls = *p.NSDLClient.AppDtls
	baseModel.Devicedtls.Deviceid = details.DeviceId
	baseModel.Deviceidentifier = models.DeviceIdentifier{
		DeviceId:    details.DeviceId,
		DeviceUnqId: details.DeviceUniqueId,
		CustUnqId:   details.CustomerUniqueId,
	}

	baseModel.Type = "ESIGNOTP"
	baseModel.Subtype = "OTPAUTHENTICATION"

	baseModel.Verstatusdtl.Verunqid = values.OtpRequiredId
	baseModel.Verstatusdtl.Smsactdtl.Smsto = userDetails.MobileCountryCode + userDetails.Msisdn
	baseModel.Verstatusdtl.Smsactdtl.Smsval = ""
	baseModel.Verstatusdtl.Smsactdtl.Smsactstatus = ""

	baseModel.Verstatusdtl.Otpdtl.Otplength = string(len(otp.OTP))
	baseModel.Verstatusdtl.Otpdtl.Otpval = otp.OTP
	baseModel.Verstatusdtl.Otpdtl.Otpencval = ""

	baseModel.Token = constants.DefaultTokenValue

	response, err := p.NSDLClient.SendPostRequest(constants.SMSVerifyEndpoint, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	var data models.SMSVerifyAPIResponse

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	if data.Respcode != constants.DefaultSuccessResponseCode {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, fmt.Errorf(data.Response)
	}

	return response.StatusCode, data, nil
}
