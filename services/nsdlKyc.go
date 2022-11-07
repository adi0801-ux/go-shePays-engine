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
	"strings"
)

func (p *ServiceConfig) CheckPANExists(kycPAN *models.KYCPAN) (int, interface{}, error) {

	//get device Id , uniqueId from db
	details, err := p.DeviceDetailsRepo.ReadDeviceDetails(kycPAN.UserID)
	if err != nil {
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
		return http.StatusBadRequest, nil, err
	}

	var data models.CustomerDeviceApiResponse
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	if data.Respcode != constants.DefaultErrorResponseCode {
		return http.StatusBadRequest, nil, fmt.Errorf(data.Response)
	}

	response, err = p.NSDLClient.SendPostRequest(constants.PANVerifyEndpoint, &baseModel)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	var dataVerify models.PANVerifyAPIResponse
	err = json.NewDecoder(response.Body).Decode(&dataVerify)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	if dataVerify.Respcode != constants.DefaultSuccessResponseCode {
		return http.StatusBadRequest, nil, fmt.Errorf(dataVerify.Response)
	}

	//save pan number
	err = p.UserKycRepo.CreateKycUserDoc(kycPAN)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	//send consent api
	StatusCode, responseBody, err := p.CustomerConsentAsking(details)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return StatusCode, responseBody, err

}

func (p *ServiceConfig) VerifyAadhar(kycAadhar *models.KYCAadharVerify) (int, interface{}, error) {

	//get device Id , uniqueId from db
	details, err := p.DeviceDetailsRepo.ReadDeviceDetails(kycAadhar.UserId)
	if err != nil {
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
		return http.StatusBadRequest, nil, err
	}
	_, err = io.Copy(filePart, file)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	_ = writer.WriteField("password", kycAadhar.Password)
	_ = writer.WriteField("token", constants.DefaultTokenValue)
	err = writer.Close()
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	header := writer.FormDataContentType()

	response, err := p.NSDLClient.SendPostFormRequest(constants.VerifyAadharEndpoint, payload, header)

	var data interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
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
		return http.StatusBadRequest, nil, err
	}

	var data models.NameValidationAPIResponse
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	if data.Respcode != constants.DefaultSuccessResponseCode {
		return http.StatusBadRequest, nil, fmt.Errorf(data.Response)
	}

	return response.StatusCode, data, nil
}

func (p *ServiceConfig) VerifySelfie(selfie *models.Selfie) (int, interface{}, error) {
	//get device Id , uniqueId from db
	details, err := p.DeviceDetailsRepo.ReadDeviceDetails(selfie.UserId)
	if err != nil {
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

	response, err := p.NSDLClient.SendPostRequest(constants.VerifySelfieEndpoint, &baseModel)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	var data models.VerifySelfieAPIResponse
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	if data.Respcode != constants.DefaultSuccessResponseCode {
		return http.StatusBadRequest, nil, fmt.Errorf(data.Response)
	}

	return response.StatusCode, data, nil
}

func (p *ServiceConfig) ValidateDOB(dob *models.ValidateDOB) (int, interface{}, error) {
	//get device Id , uniqueId from db
	details, err := p.DeviceDetailsRepo.ReadDeviceDetails(dob.UserId)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	additionalDetails, err := p.CustomerDetailsRepo.ReadCustomerAdditionalDetails(dob.UserId)
	if err != nil {
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
		return http.StatusBadRequest, nil, err
	}

	var data models.ValidateDOBAPIResponse
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	if data.Respcode != constants.DefaultSuccessResponseCode {
		return http.StatusBadRequest, nil, fmt.Errorf(data.Response)
	}

	return response.StatusCode, data, nil
}

func (p *ServiceConfig) ValidatePAN(details *models.DeviceDetails) (int, interface{}, error) {

	kycDoc, err := p.UserKycRepo.ReadKycUserDoc(details.UserID)
	if err != nil {
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
		return http.StatusBadRequest, nil, err
	}

	var data models.ValidatePANAPIResponse
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	if data.Respcode != constants.DefaultSuccessResponseCode {
		return http.StatusBadRequest, nil, fmt.Errorf(data.Response)
	}

	return response.StatusCode, data, nil
}
