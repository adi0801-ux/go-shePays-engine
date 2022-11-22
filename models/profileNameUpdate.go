package models

type ProfileNameUpdateAPI struct {
	Channelid  string  `json:"channelid"`
	Appdtls    AppDtls `json:"appdtls"`
	Devicedtls struct {
		Deviceid string `json:"deviceid"`
	} `json:"devicedtls"`
	Deviceidentifier DeviceIdentifier `json:"deviceidentifier"`
	FirstName        string           `json:"first_name"`
	MiddleName       string           `json:"middle_name"`
	LastName         string           `json:"last_name"`
	Token            string           `json:"token"`
	Signcs           string           `json:"signcs"`
}
