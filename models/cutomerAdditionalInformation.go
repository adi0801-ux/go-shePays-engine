package models

type CustomerAdditionalInformationAPI struct {
	Channelid  string  `json:"channelid"`
	Appdtls    AppDtls `json:"appdtls"`
	Devicedtls struct {
		Deviceid string `json:"deviceid"`
	} `json:"devicedtls"`
	Deviceidentifier DeviceIdentifier `json:"deviceidentifier"`
	Custmindtls      struct {
		Customername string `json:"customername"`
		Custemail    string `json:"custemail"`
		Dateofbirth  string `json:"dateofbirth"`
		Gender       string `json:"gender"`
		Salutation   string `json:"salutation"`
		Country      string `json:"country"`
		Pincode      string `json:"pincode"`
		State        string `json:"state"`
		City         string `json:"city"`
		Msisdn       string `json:"msisdn"`
		Mobcntrycode string `json:"mobcntrycode"`
	} `json:"custmindtls"`
	Token  string `json:"token"`
	Signcs string `json:"signcs"`
}

type CustomerAdditionalInformationAPIResponse struct {
	Response string `json:"response"`
	Respcode string `json:"respcode"`
}
