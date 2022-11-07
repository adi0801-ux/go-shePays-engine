package models

type ConsentAskingAPI struct {
	Channelid  string  `json:"channelid"`
	Appdtls    AppDtls `json:"appdtls"`
	Devicedtls struct {
		Deviceid string `json:"deviceid"`
	} `json:"devicedtls"`
	Deviceidentifier DeviceIdentifier `json:"deviceidentifier"`
	Consentflg       string           `json:"consentflg"`
	Token            string           `json:"token"`
	Signcs           string           `json:"signcs"`
}

type ConsentAskingAPIResponse struct {
	Response string `json:"response"`
	Respcode string `json:"respcode"`
}
