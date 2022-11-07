package models

type BaseNsdl struct {
	Channelid        string           `json:"channelid"`
	Appdtls          AppDtls          `json:"appdtls"`
	Deviceidentifier DeviceIdentifier `json:"deviceidentifier"`
	Devicedtls       struct {
		Deviceid string `json:"deviceid"`
	} `json:"devicedtls"`
}

type AppDtls struct {
	Appid       string `json:"appid"`
	ApplVersion string `json:"applversion"`
	AppRegFlg   string `json:"appregflg"`
	Pushnkey    string `json:"pushnkey"`
}

type DeviceIdentifier struct {
	DeviceId    string `json:"deviceid"`
	DeviceUnqId string `json:"deviceunqid"`
	CustUnqId   string `json:"custunqid"`
}
