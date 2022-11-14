package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"shepays/constants"
	"shepays/models"
	"shepays/utils"
	"strings"
)

func (p *ServiceConfig) CheckPANExists(kycPAN *models.KYCPAN) (int, interface{}, error) {

	//get device Id , uniqueId from db
	details, err := p.DeviceDetailsRepo.ReadDeviceDetails(kycPAN.UserID)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	baseModel := models.PANCheckExitsAPI{}
	baseModel.Channelid = p.NSDLClient.ChannelId
	baseModel.Appdtls = *p.NSDLClient.AppDtls
	baseModel.Devicedtls.Deviceid = details.DeviceId
	baseModel.Deviceidentifier = models.DeviceIdentifier{
		DeviceId:    details.DeviceId,
		DeviceUnqId: details.DeviceUniqueId,
		CustUnqId:   details.CustomerUniqueId,
	}
	baseModel.Panno = kycPAN.PanNumber

	baseModel.Token = constants.DefaultTokenValue

	response, err := p.NSDLClient.SendPostRequest(constants.CheckPanExistsEndpoint, &baseModel)
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
	if data.Respcode != constants.DefaultErrorResponseCode {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, fmt.Errorf(data.Response)
	}

	response, err = p.NSDLClient.SendPostRequest(constants.PANVerifyEndpoint, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	var dataVerify models.PANVerifyAPIResponse
	err = json.NewDecoder(response.Body).Decode(&dataVerify)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	if dataVerify.Respcode != constants.DefaultSuccessResponseCode {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, fmt.Errorf(dataVerify.Response)
	}

	//save pan number
	err = p.UserKycRepo.CreateKycUserDoc(kycPAN)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	//send consent api
	StatusCode, responseBody, err := p.CustomerConsentAsking(details)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	return StatusCode, responseBody, err

}

func (p *ServiceConfig) VerifyAadhar(kycAadhar *models.KYCAadharVerify) (int, interface{}, error) {

	//get device Id , uniqueId from db
	details, err := p.DeviceDetailsRepo.ReadDeviceDetails(kycAadhar.UserId)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	payload := new(bytes.Buffer)
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("channelid", p.NSDLClient.ChannelId)
	_ = writer.WriteField("appdtls.appid", p.NSDLClient.AppDtls.Appid)
	_ = writer.WriteField("appdtls.applversion", p.NSDLClient.AppDtls.ApplVersion)
	_ = writer.WriteField("appdtls.appregflg", p.NSDLClient.AppDtls.AppRegFlg)
	_ = writer.WriteField("appdtls.pushnkey", "")
	_ = writer.WriteField("deviceidentifier.deviceid", details.DeviceId)
	_ = writer.WriteField("deviceidentifier.deviceunqid", details.DeviceUniqueId)
	_ = writer.WriteField("deviceidentifier.custunqid", details.CustomerUniqueId)

	file, err := kycAadhar.AddharFile.Open()
	defer file.Close()

	filePart, errFile := writer.CreateFormFile("aadharFile", kycAadhar.AddharFile.Filename)
	if errFile != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	_, err = io.Copy(filePart, file)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	_ = writer.WriteField("password", kycAadhar.Password)
	_ = writer.WriteField("token", constants.DefaultTokenValue)
	err = writer.Close()
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	header := writer.FormDataContentType()

	response, err := p.NSDLClient.SendPostFormRequest(constants.VerifyAadharEndpoint, payload, header)

	var data interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	//name match api
	StatusCode, dataVerify, err := p.ValidateName(details)
	if err != nil {
		return StatusCode, dataVerify, err
	}

	return response.StatusCode, nil, err

}

func (p *ServiceConfig) ValidateName(details *models.DeviceDetails) (int, interface{}, error) {

	baseModel := models.BaseNsdl{}
	baseModel.Channelid = p.NSDLClient.ChannelId
	baseModel.Appdtls = *p.NSDLClient.AppDtls
	baseModel.Devicedtls.Deviceid = details.DeviceId
	baseModel.Deviceidentifier = models.DeviceIdentifier{
		DeviceId:    details.DeviceId,
		DeviceUnqId: details.DeviceUniqueId,
		CustUnqId:   details.CustomerUniqueId,
	}

	response, err := p.NSDLClient.SendPostRequest(constants.ValidateNameEndpoint, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	var data models.NameValidationAPIResponse
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

func (p *ServiceConfig) VerifySelfie(selfie *models.Selfie) (int, interface{}, error) {
	//get device Id , uniqueId from db
	details, err := p.DeviceDetailsRepo.ReadDeviceDetails(selfie.UserId)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	baseModel := models.VerifySelfieAPI{}
	baseModel.Channelid = p.NSDLClient.ChannelId
	baseModel.Appdtls = *p.NSDLClient.AppDtls
	baseModel.Devicedtls.Deviceid = details.DeviceId
	baseModel.Deviceidentifier = models.DeviceIdentifier{
		DeviceId:    details.DeviceId,
		DeviceUnqId: details.DeviceUniqueId,
		CustUnqId:   details.CustomerUniqueId,
	}

	//need to fill rest
	baseModel.SelfieImage = selfie.Image
	baseModel.Referid = selfie.ReferId
	baseModel.Jobid = selfie.JobId

	response, err := p.NSDLClient.SendPostRequest(constants.VerifySelfieEndpoint, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	var data models.VerifySelfieAPIResponse
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	if data.Facedtl.Facematch == "1" {
		utils.Log.Error(err)
		return http.StatusBadRequest, data, fmt.Errorf("face recoginition failed")
	}

	if data.Respcode != constants.DefaultSuccessResponseCode {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, fmt.Errorf(data.Response)
	}

	return response.StatusCode, data, nil
}

func (p *ServiceConfig) ValidateDOB(dob *models.ValidateDOB) (int, interface{}, error) {
	//get device Id , uniqueId from db
	details, err := p.DeviceDetailsRepo.ReadDeviceDetails(dob.UserId)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	additionalDetails, err := p.CustomerDetailsRepo.ReadCustomerAdditionalDetails(dob.UserId)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	dobString := strings.Split(additionalDetails.DOB, "-")

	baseModel := models.ValidateDOBAPI{}
	baseModel.Channelid = p.NSDLClient.ChannelId
	baseModel.Appdtls = *p.NSDLClient.AppDtls
	baseModel.Devicedtls.Deviceid = details.DeviceId
	baseModel.Deviceidentifier = models.DeviceIdentifier{
		DeviceId:    details.DeviceId,
		DeviceUnqId: details.DeviceUniqueId,
		CustUnqId:   details.CustomerUniqueId,
	}

	baseModel.Objective = dob.Objective
	baseModel.ReferenceNo = dob.ReferenceNo
	baseModel.Dateofbirth = dobString[2] + "-" + dobString[1] + "-" + dobString[0]
	//baseModel.Dateofbirth = additionalDetails.DOB
	//need to fill rest

	response, err := p.NSDLClient.SendPostRequest(constants.ValidateDOBEndpoint, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	var data models.ValidateDOBAPIResponse
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	if data.Respcode != constants.DefaultSuccessResponseCode {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, fmt.Errorf(data.Response)
	}

	StatusCode, dataVerify, err := p.ValidatePAN(details)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	return StatusCode, dataVerify, nil
}

func (p *ServiceConfig) ValidatePAN(details *models.DeviceDetails) (int, interface{}, error) {

	kycDoc, err := p.UserKycRepo.ReadKycUserDoc(details.UserID)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	baseModel := models.ValidatePANAPI{}
	baseModel.Channelid = p.NSDLClient.ChannelId
	baseModel.Appdtls = *p.NSDLClient.AppDtls
	baseModel.Devicedtls.Deviceid = details.DeviceId
	baseModel.Deviceidentifier = models.DeviceIdentifier{
		DeviceId:    details.DeviceId,
		DeviceUnqId: details.DeviceUniqueId,
		CustUnqId:   details.CustomerUniqueId,
	}

	//need to fill rest
	baseModel.Panno = kycDoc.PanNumber

	response, err := p.NSDLClient.SendPostRequest(constants.ValidatePANEndpoint, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	var data models.ValidatePANAPIResponse
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

func (p *ServiceConfig) GetProductCardDetails(details *models.DeviceDetails) (int, interface{}, error) {

	baseModel := models.GetCardDetailsAPI{}
	baseModel.Channelid = p.NSDLClient.ChannelId
	baseModel.Appdtls = *p.NSDLClient.AppDtls
	baseModel.Devicedtls.Deviceid = details.DeviceId
	baseModel.Deviceidentifier = models.DeviceIdentifier{
		DeviceId:    details.DeviceId,
		DeviceUnqId: details.DeviceUniqueId,
		CustUnqId:   details.CustomerUniqueId,
	}

	response, err := p.NSDLClient.SendPostRequest(constants.GetCardProductsEndpoint, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	var data models.GetCardDetailsAPIResponse
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	if data.Respcode != constants.DefaultSuccessResponseCode {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, fmt.Errorf(data.Response)
	}

	userCardInformation := models.UserCardCreateInformation{}

	userCardInformation.AccountProdCode = data.AccntProdlist[0].AccntProdCode
	userCardInformation.AccountProdName = data.AccntProdlist[0].AccntProdName

	for _, cards := range data.CardProdlist {

		if cards.Network == constants.DefaultCardNetwork {
			userCardInformation.CardNetwork = cards.Network
			userCardInformation.CardProdName = cards.CardProdName
			userCardInformation.CardProdCode = cards.CardProdCode
			userCardInformation.BinPrefix = cards.BinPrifix
		}
	}

	userCardInformation.UserID = details.UserID

	err = p.UserCardInformationRepo.CreateUserCardCreateInformation(&userCardInformation)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	return response.StatusCode, userCardInformation, nil

}

func (p *ServiceConfig) AoFAPI(aof *models.AoFModel) (int, interface{}, error) {

	//get device Id , uniqueId from db
	details, err := p.DeviceDetailsRepo.ReadDeviceDetails(aof.UserId)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	additionalInformation, err := p.CustomerDetailsRepo.ReadCustomerAdditionalDetails(details.UserID)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	userDetails, err := p.CustomerDetailsRepo.ReadCustomerDetails(details.UserID)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	_, cardInfo, err := p.GetProductCardDetails(details)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	userKyc, err := p.UserKycRepo.ReadKycUserDoc(details.UserID)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	userCardInfo := cardInfo.(models.UserCardCreateInformation)

	dobString := strings.Split(additionalInformation.DOB, "-")

	baseModel := models.AoFModelAPI{}
	baseModel.Channelid = p.NSDLClient.ChannelId
	baseModel.Appdtls = *p.NSDLClient.AppDtls
	baseModel.Devicedtls.Deviceid = details.DeviceId
	baseModel.Deviceidentifier = models.DeviceIdentifier{
		DeviceId:    details.DeviceId,
		DeviceUnqId: details.DeviceUniqueId,
		CustUnqId:   details.CustomerUniqueId,
	}

	baseModel.Custcreddtls.Userid = userDetails.MobileCountryCode + userDetails.Msisdn
	baseModel.Custcreddtls.Usertype = "CUSTOMER"
	baseModel.Custcreddtls.Role = "CUST"

	baseModel.Aofdtls.Prefix = additionalInformation.Salutation
	baseModel.Aofdtls.FirstName = userDetails.CustomerFName
	baseModel.Aofdtls.MidName = userDetails.CustomerMName
	baseModel.Aofdtls.LastName = userDetails.CustomerLName
	baseModel.Aofdtls.MotherMaidenName = ""
	baseModel.Aofdtls.FatherName = additionalInformation.FathersName
	baseModel.Aofdtls.MaritalStatus = additionalInformation.MartialStatus
	baseModel.Aofdtls.CustomerMobilePhone = userDetails.MobileCountryCode + userDetails.Msisdn
	baseModel.Aofdtls.DateOfBirth = dobString[2] + dobString[1] + dobString[0]
	baseModel.Aofdtls.IncomeCategory = additionalInformation.IncomeCategory
	baseModel.Aofdtls.EmailId = userDetails.EmailId

	//unknown
	baseModel.Aofdtls.NationalIdentificationCode = userKyc.PanNumber

	baseModel.Aofdtls.ProfessionCode = additionalInformation.ProfessionalCode
	baseModel.Aofdtls.Relation = additionalInformation.Relation
	baseModel.Aofdtls.Sex = userDetails.Gender
	baseModel.Aofdtls.ProductCodeString = aof.ProductCodeString

	baseModel.Aofdtls.Address.City = additionalInformation.City
	baseModel.Aofdtls.Address.Country = additionalInformation.Country
	baseModel.Aofdtls.Address.Line1 = additionalInformation.AddressLine1
	baseModel.Aofdtls.Address.Line2 = additionalInformation.AddressLine2
	baseModel.Aofdtls.Address.Line3 = "__ None __"
	baseModel.Aofdtls.Address.State = additionalInformation.State
	baseModel.Aofdtls.Address.Pincode = additionalInformation.PinCode

	baseModel.Aofdtls.AccountTitle = userDetails.CustomerFName + userDetails.CustomerMName + userDetails.CustomerLName
	baseModel.Aofdtls.AcctCurrencyString = "1"
	baseModel.Aofdtls.BranchCode = "8888"
	baseModel.Aofdtls.CustomerIDString = ""
	baseModel.Aofdtls.FlgJointHolderString = "N"
	baseModel.Aofdtls.FlgRestrictAcctString = "N"
	baseModel.Aofdtls.FlgSCWaiveString = "N"
	baseModel.Aofdtls.FlgTransactionType = "A"
	baseModel.Aofdtls.MinorAcctStatusString = "N"
	baseModel.Aofdtls.Nationality = additionalInformation.Country
	baseModel.Aofdtls.CustomerEducation = additionalInformation.CustomerEducation
	baseModel.Aofdtls.IsStaff = "N"
	baseModel.Aofdtls.IcType = "I"
	baseModel.Aofdtls.Origin = "IND"

	baseModel.Aofdtls.Name.FirstName = userDetails.CustomerFName
	baseModel.Aofdtls.Name.FormattedFullName = userDetails.CustomerName
	baseModel.Aofdtls.Name.FullName = userDetails.CustomerName
	baseModel.Aofdtls.Name.LastName = userDetails.CustomerLName
	baseModel.Aofdtls.Name.MidName = userDetails.CustomerMName
	baseModel.Aofdtls.Name.Prefix = additionalInformation.Salutation
	baseModel.Aofdtls.Name.ShortName = userDetails.CustomerNickname
	baseModel.Aofdtls.Name.SingleFullName = userDetails.CustomerFName + userDetails.CustomerMName + userDetails.CustomerLName

	//save this information
	baseModel.Carddtls.CardProdCode = userCardInfo.CardProdCode
	baseModel.Carddtls.CardProdName = userCardInfo.CardProdName
	baseModel.Carddtls.Network = userCardInfo.CardNetwork

	response, err := p.NSDLClient.SendPostRequest(constants.AOFCreationEndpoint, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	var data models.AoFModelAPIResponse
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	if data.Respcode != constants.DefaultSuccessResponseCode {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, fmt.Errorf(data.Response)
	}

	//save refnumber
	interm := models.UserIntermValues{
		ReferenceNumber: data.Refno,
		UserID:          details.UserID,
	}
	err = p.IntermValuesRepo.CreateOrUpdateUserIntermValues(&interm)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	return response.StatusCode, data, nil

}

func (p *ServiceConfig) VCifAPI(user *models.UserId) (int, interface{}, error) {

	//get device Id , uniqueId from db
	details, err := p.DeviceDetailsRepo.ReadDeviceDetails(user.UserId)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	additionalInformation, err := p.CustomerDetailsRepo.ReadCustomerAdditionalDetails(details.UserID)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	userDetails, err := p.CustomerDetailsRepo.ReadCustomerDetails(details.UserID)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	userKyc, err := p.UserKycRepo.ReadKycUserDoc(details.UserID)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	dobString := strings.Split(additionalInformation.DOB, "-")

	baseModel := models.VCifAPI{}
	baseModel.Channelid = p.NSDLClient.ChannelId
	baseModel.Appdtls = *p.NSDLClient.AppDtls
	baseModel.Devicedtls.Deviceid = details.DeviceId
	baseModel.Deviceidentifier = models.DeviceIdentifier{
		DeviceId:    details.DeviceId,
		DeviceUnqId: details.DeviceUniqueId,
		CustUnqId:   details.CustomerUniqueId,
	}

	baseModel.Custcreddtls.Userid = userDetails.MobileCountryCode + userDetails.Msisdn
	baseModel.Custcreddtls.Usertype = "CUSTOMER"
	baseModel.Custcreddtls.Role = "CUST"

	baseModel.Custcreddtls.FirstName = userDetails.CustomerFName
	baseModel.Custcreddtls.MiddleName = userDetails.CustomerMName
	baseModel.Custcreddtls.LastName = userDetails.CustomerLName
	baseModel.Custcreddtls.FullName = userDetails.CustomerName

	baseModel.Custcreddtls.Mobileno = userDetails.MobileCountryCode + userDetails.Msisdn
	baseModel.Custcreddtls.Mobcountrycode = userDetails.MobileCountryCode
	baseModel.Custcreddtls.Emailid = userDetails.EmailId
	baseModel.Custcreddtls.Catcode = "NA"

	baseModel.Createindvcifdtl.ExtCustomerID = ""

	baseModel.Createindvcifdtl.IndividualCustomerDTO.Sex = userDetails.Gender
	baseModel.Createindvcifdtl.IndividualCustomerDTO.MaritalStatus = additionalInformation.MartialStatus
	baseModel.Createindvcifdtl.IndividualCustomerDTO.ProfessionCode = additionalInformation.ProfessionalCode
	baseModel.Createindvcifdtl.IndividualCustomerDTO.IsStaff = "N"
	baseModel.Createindvcifdtl.IndividualCustomerDTO.CustomerId = ""
	baseModel.Createindvcifdtl.IndividualCustomerDTO.EmployeeId = "1"
	baseModel.Createindvcifdtl.IndividualCustomerDTO.CustomerEducation = additionalInformation.CustomerEducation
	baseModel.Createindvcifdtl.IndividualCustomerDTO.MotherMaidenName = ""
	baseModel.Createindvcifdtl.IndividualCustomerDTO.CountryOfResidence = "IN"
	baseModel.Createindvcifdtl.IndividualCustomerDTO.CustomerMobilePhone = userDetails.MobileCountryCode + userDetails.Msisdn
	baseModel.Createindvcifdtl.IndividualCustomerDTO.EmailId = userDetails.EmailId
	baseModel.Createindvcifdtl.IndividualCustomerDTO.Nationality = "IND"

	//unknown field
	baseModel.Createindvcifdtl.IndividualCustomerDTO.NationalIdentificationCode = userKyc.PanNumber

	baseModel.Createindvcifdtl.IndividualCustomerDTO.IcType = "I"
	baseModel.Createindvcifdtl.IndividualCustomerDTO.Category = "GI"

	baseModel.Createindvcifdtl.IndividualCustomerDTO.Name.FirstName = userDetails.CustomerFName
	baseModel.Createindvcifdtl.IndividualCustomerDTO.Name.FormattedFullName = userDetails.CustomerName
	baseModel.Createindvcifdtl.IndividualCustomerDTO.Name.FullName = userDetails.CustomerName
	baseModel.Createindvcifdtl.IndividualCustomerDTO.Name.LastName = userDetails.CustomerLName
	baseModel.Createindvcifdtl.IndividualCustomerDTO.Name.MidName = userDetails.CustomerMName
	baseModel.Createindvcifdtl.IndividualCustomerDTO.Name.Prefix = additionalInformation.Salutation
	baseModel.Createindvcifdtl.IndividualCustomerDTO.Name.ShortName = userDetails.CustomerNickname
	baseModel.Createindvcifdtl.IndividualCustomerDTO.Name.SingleFullName = userDetails.CustomerFName + userDetails.CustomerMName + userDetails.CustomerLName

	baseModel.Createindvcifdtl.IndividualCustomerDTO.DateOfBirthOrRegistration = dobString[2] + dobString[1] + dobString[0]
	baseModel.Createindvcifdtl.IndividualCustomerDTO.SignatureType = "1"
	baseModel.Createindvcifdtl.IndividualCustomerDTO.Phone = userDetails.MobileCountryCode + userDetails.Msisdn
	baseModel.Createindvcifdtl.IndividualCustomerDTO.Origin = "IND"
	baseModel.Createindvcifdtl.IndividualCustomerDTO.Language = "eng"
	baseModel.Createindvcifdtl.IndividualCustomerDTO.HomeBranchCode = "8888"
	baseModel.Createindvcifdtl.IndividualCustomerDTO.CifType = "C"

	baseModel.Createindvcifdtl.IndividualCustomerDTO.Address.City = additionalInformation.City
	baseModel.Createindvcifdtl.IndividualCustomerDTO.Address.Country = additionalInformation.Country
	baseModel.Createindvcifdtl.IndividualCustomerDTO.Address.Line1 = additionalInformation.AddressLine1
	baseModel.Createindvcifdtl.IndividualCustomerDTO.Address.Line2 = additionalInformation.AddressLine2
	baseModel.Createindvcifdtl.IndividualCustomerDTO.Address.Line3 = "__ None __"
	baseModel.Createindvcifdtl.IndividualCustomerDTO.Address.State = additionalInformation.State
	baseModel.Createindvcifdtl.IndividualCustomerDTO.Address.Pincode = additionalInformation.PinCode

	baseModel.Createindvcifdtl.IndividualCustomerDTO.PermanantAddress.City = additionalInformation.City
	baseModel.Createindvcifdtl.IndividualCustomerDTO.PermanantAddress.Country = additionalInformation.Country
	baseModel.Createindvcifdtl.IndividualCustomerDTO.PermanantAddress.Line1 = additionalInformation.AddressLine1
	baseModel.Createindvcifdtl.IndividualCustomerDTO.PermanantAddress.Line2 = additionalInformation.AddressLine2
	baseModel.Createindvcifdtl.IndividualCustomerDTO.PermanantAddress.Line3 = "__ None __"
	baseModel.Createindvcifdtl.IndividualCustomerDTO.PermanantAddress.State = additionalInformation.State
	baseModel.Createindvcifdtl.IndividualCustomerDTO.PermanantAddress.Pincode = additionalInformation.PinCode

	baseModel.Createindvcifdtl.IndividualCustomerDTO.Relation = additionalInformation.Relation
	baseModel.Createindvcifdtl.IndividualCustomerDTO.FatherName = additionalInformation.FathersName

	baseModel.Createindvcifdtl.MisClass = "DIVISION"
	baseModel.Createindvcifdtl.IsFCPB = "true"
	baseModel.Createindvcifdtl.MisCode = "DIVISION"
	baseModel.Createindvcifdtl.IncomeCategory = additionalInformation.IncomeCategory

	response, err := p.NSDLClient.SendPostRequest(constants.VCIFCreationEndpoint, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	var data models.VCifAPIResponse
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

func (p *ServiceConfig) AccountCreateProxy(user *models.UserId) (int, interface{}, error) {

	//get device Id , uniqueId from db
	details, err := p.DeviceDetailsRepo.ReadDeviceDetails(user.UserId)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	userDetails, err := p.CustomerDetailsRepo.ReadCustomerDetails(details.UserID)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	userCardInfo, err := p.UserCardInformationRepo.ReadUserCardCreateInformation(details.UserID)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	baseModel := models.AccountCreationAPI{}
	baseModel.Channelid = p.NSDLClient.ChannelId
	baseModel.Appdtls = *p.NSDLClient.AppDtls
	baseModel.Devicedtls.Deviceid = details.DeviceId
	baseModel.Deviceidentifier = models.DeviceIdentifier{
		DeviceId:    details.DeviceId,
		DeviceUnqId: details.DeviceUniqueId,
		CustUnqId:   details.CustomerUniqueId,
	}

	baseModel.Custcreddtls.Userid = userDetails.MobileCountryCode + userDetails.Msisdn
	baseModel.Custcreddtls.Usertype = "CUSTOMER"
	baseModel.Custcreddtls.Role = "CUST"

	//unknown
	baseModel.Custcreddtls.Cif = ""

	baseModel.Accountcreationdtl.AccountNoString = ""
	baseModel.Accountcreationdtl.AccountTitle = userDetails.CustomerFName + userDetails.CustomerMName + userDetails.CustomerLName
	baseModel.Accountcreationdtl.AcctCurrencyString = "1"
	baseModel.Accountcreationdtl.BranchCode = "8888"

	//unknown
	baseModel.Accountcreationdtl.CustomerIDString = "N"

	baseModel.Accountcreationdtl.FlgJointHolderString = "N"
	baseModel.Accountcreationdtl.FlgRestrictAcctString = "N"
	baseModel.Accountcreationdtl.FlgSCWaiveString = "N"
	baseModel.Accountcreationdtl.FlgTransactionType = "A"
	baseModel.Accountcreationdtl.MinorAcctStatusString = "N"

	//unknown
	baseModel.Accountcreationdtl.ProductCodeString = userCardInfo.AccountProdCode

	response, err := p.NSDLClient.SendPostRequest(constants.AccountCreationEndpoint, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	var data models.AccountCreationAPIResponse
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	if data.Respcode != constants.DefaultSuccessResponseCode {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, fmt.Errorf(data.Response)
	}

	//save account number
	account := &models.UserAccount{
		AccountNumber: data.AccDtls.Accountno,
		UserID:        details.UserID,
	}
	err = p.UserAccountRepo.CreateUserAccount(account)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	return p.CreateCardProxy(user)

}

func (p *ServiceConfig) CreateCardProxy(user *models.UserId) (int, interface{}, error) {

	//get device Id , uniqueId from db
	details, err := p.DeviceDetailsRepo.ReadDeviceDetails(user.UserId)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	additionalInformation, err := p.CustomerDetailsRepo.ReadCustomerAdditionalDetails(details.UserID)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	userCardInfo, err := p.UserCardInformationRepo.ReadUserCardCreateInformation(details.UserID)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	baseModel := models.CardCreationAPI{}
	baseModel.Channelid = p.NSDLClient.ChannelId
	baseModel.Appdtls = *p.NSDLClient.AppDtls
	baseModel.Devicedtls.Deviceid = details.DeviceId
	baseModel.Deviceidentifier = models.DeviceIdentifier{
		DeviceId:    details.DeviceId,
		DeviceUnqId: details.DeviceUniqueId,
		CustUnqId:   details.CustomerUniqueId,
	}

	account, err := p.UserAccountRepo.ReadUserAccount(details.UserID)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	//fill from saved data
	baseModel.Msg.AccountNo = account.AccountNumber

	baseModel.Msg.NetworkType = userCardInfo.CardNetwork

	baseModel.Msg.BinPrefix = userCardInfo.BinPrefix

	baseModel.Msg.AddressDtls.City = additionalInformation.City
	baseModel.Msg.AddressDtls.Country = additionalInformation.Country
	baseModel.Msg.AddressDtls.Address1 = additionalInformation.AddressLine1
	baseModel.Msg.AddressDtls.Address2 = additionalInformation.AddressLine2
	baseModel.Msg.AddressDtls.Address3 = "__ None __"
	baseModel.Msg.AddressDtls.State = additionalInformation.State
	baseModel.Msg.AddressDtls.Pincode = additionalInformation.PinCode

	baseModel.Msg.SessionId = ""

	response, err := p.NSDLClient.SendPostRequest(constants.CardCreationEndpoint, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	var data models.AccountCreationAPIResponse
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
