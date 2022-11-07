package models

import "time"

// register device
type RegisterDevice struct {
	Os               string `json:"os"`
	OsVersion        string `json:"os_version"`
	DeviceMake       string `json:"device_make"`
	DeviceModel      string `json:"device_model"`
	DeviceProduct    string `json:"device_product"`
	DeviceSerial     string `json:"device_serial"`
	DeviceBrand      string `json:"device_brand"`
	DeviceSdkVersion string `json:"device_sdk_version"`
	RootFlag         string `json:"root_flg"`
	SimFlag          string `json:"sim_flag"`
	LongDetail       string `json:"long_detail"`
	LatDetail        string `json:"lat_detail"`
	LocationDetail   string `json:"location_detail"`
	IpAddress        string `json:"ip_address"`
	UserId           string `json:"user_id"`
}

type RegiterDeviceApi struct {
	Mbappdetails struct {
		Channelid  string  `json:"channelid"`
		Appdtls    AppDtls `json:"appdtls"`
		Devicedtls struct {
			Deviceid         string `json:"deviceid"`
			Os               string `json:"os"`
			OsVersion        string `json:"osVersion"`
			Devicemake       string `json:"devicemake"`
			Devicemodel      string `json:"devicemodel"`
			Deviceproduct    string `json:"deviceproduct"`
			Deviceserial     string `json:"deviceserial"`
			Devicebrand      string `json:"devicebrand"`
			Devicesdkversion string `json:"devicesdkversion"`
			Rootflg          string `json:"rootflg"`
			Simflag          string `json:"simflag"`
		} `json:"devicedtls"`
		Locationdtls struct {
			Longdtl     string `json:"longdtl"`
			Latdtl      string `json:"latdtl"`
			LocationDtl string `json:"locationDtl"`
			IpAdress    string `json:"ipAdress"`
		} `json:"locationdtls"`
	} `json:"mbappdetails"`
	Token  string `json:"token"`
	Signcs string `json:"signcs"`
}

type RegisterDeviceResponse struct {
	Deviceidentifier struct {
		Deviceunqid string `json:"deviceunqid"`
	} `json:"deviceidentifier"`
	Response string `json:"response"`
	Respcode string `json:"respcode"`
}

const TableNameDeviceDetails = "device_details"

type DeviceDetails struct {
	DeviceId         string    `gorm:"column:device_id;not null" json:"device_id"`
	DeviceUniqueId   string    `gorm:"column:device_unique_id;not null" json:"device_unique_id"`
	CustomerUniqueId string    `gorm:"column:customer_unique_id" json:"customer_unique_id"`
	UserID           string    `gorm:"column:user_id;not null" json:"user_id"`
	ID               int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt        time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
}

func (*DeviceDetails) TableName() string {
	return TableNameDeviceDetails
}
