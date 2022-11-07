package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shepays/constants"
	"shepays/models"
	"shepays/utils"
)

func (p *ServiceConfig) CheckDeviceVersion(deviceIdentifier *models.DeviceIdentifier) (int, interface{}, error) {

	baseModel := models.BaseNsdl{
		Channelid:        p.NSDLClient.ChannelId,
		Appdtls:          *p.NSDLClient.AppDtls,
		Deviceidentifier: *deviceIdentifier,
	}

	response, err := p.NSDLClient.SendPostRequest(constants.VersionCheckEndpoint, &baseModel)
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

func (p *ServiceConfig) RegisterDevice(registerDevice *models.RegisterDevice) (int, interface{}, error) {

	baseModel := models.RegiterDeviceApi{}

	baseModel.Mbappdetails.Channelid = p.NSDLClient.ChannelId
	baseModel.Mbappdetails.Appdtls = *p.NSDLClient.AppDtls

	//device details
	baseModel.Mbappdetails.Devicedtls.Deviceid = utils.GenerateDeviceId()
	baseModel.Mbappdetails.Devicedtls.Os = registerDevice.Os
	baseModel.Mbappdetails.Devicedtls.OsVersion = registerDevice.OsVersion
	baseModel.Mbappdetails.Devicedtls.Devicemake = registerDevice.DeviceMake
	baseModel.Mbappdetails.Devicedtls.Devicemodel = registerDevice.DeviceModel
	baseModel.Mbappdetails.Devicedtls.Deviceproduct = registerDevice.DeviceProduct
	baseModel.Mbappdetails.Devicedtls.Deviceserial = registerDevice.DeviceSerial
	baseModel.Mbappdetails.Devicedtls.Devicebrand = registerDevice.DeviceBrand
	baseModel.Mbappdetails.Devicedtls.Devicesdkversion = registerDevice.DeviceSdkVersion
	baseModel.Mbappdetails.Devicedtls.Rootflg = registerDevice.RootFlag
	baseModel.Mbappdetails.Devicedtls.Simflag = registerDevice.SimFlag

	//location details
	baseModel.Mbappdetails.Locationdtls.Longdtl = registerDevice.LongDetail
	baseModel.Mbappdetails.Locationdtls.Latdtl = registerDevice.LatDetail
	baseModel.Mbappdetails.Locationdtls.LocationDtl = registerDevice.LocationDetail
	baseModel.Mbappdetails.Locationdtls.IpAdress = registerDevice.IpAddress

	//token
	baseModel.Token = constants.DefaultTokenValue

	response, err := p.NSDLClient.SendPostRequest(constants.RegisterDeviceEndpoint, &baseModel)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	var data models.RegisterDeviceResponse
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	if data.Respcode != constants.DefaultSuccessResponseCode {
		return http.StatusBadRequest, nil, fmt.Errorf(data.Response)
	}

	//save to db

	//create db model
	deviceDetail := &models.DeviceDetails{
		DeviceId:       baseModel.Mbappdetails.Devicedtls.Deviceid,
		DeviceUniqueId: data.Deviceidentifier.Deviceunqid,
		UserID:         registerDevice.UserId,
	}
	err = p.DeviceDetailsRepo.CreateDeviceDetails(deviceDetail)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return response.StatusCode, map[string]string{"device_unique_id": data.Deviceidentifier.Deviceunqid, "device_id": baseModel.Mbappdetails.Devicedtls.Deviceid}, nil
}
