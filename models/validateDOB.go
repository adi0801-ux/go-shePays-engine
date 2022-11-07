package models

type ValidateDOBAPI struct {
	Channelid  string  `json:"channelid"`
	Appdtls    AppDtls `json:"appdtls"`
	Devicedtls struct {
		Deviceid string `json:"deviceid"`
	} `json:"devicedtls"`
	Deviceidentifier DeviceIdentifier `json:"deviceidentifier"`
	Objective        string           `json:"objective"`
	ReferenceNo      string           `json:"referenceNo"`
	Dateofbirth      string           `json:"dateofbirth"`
	Token            string           `json:"token"`
	Signcs           string           `json:"signcs"`
}

type ValidateDOB struct {
	Objective   string `json:"objective"`
	ReferenceNo string `json:"reference_no"`
	UserId      string `json:"user_id"`
}

type ValidateDOBAPIResponse struct {
	Reason   string `json:"reason"`
	Agemth   int    `json:"agemth"`
	Allowed  string `json:"allowed"`
	Response string `json:"response"`
	Ageyrs   int    `json:"ageyrs"`
	Respcode string `json:"respcode"`
	Ageday   int    `json:"ageday"`
}
