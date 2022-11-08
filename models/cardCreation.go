package models

type CardCreationAPI struct {
	Channelid  string  `json:"channelid"`
	Appdtls    AppDtls `json:"appdtls"`
	Devicedtls struct {
		Deviceid string `json:"deviceid"`
	} `json:"devicedtls"`
	Deviceidentifier DeviceIdentifier `json:"deviceidentifier"`
	Msg              struct {
		AccountNo   string `json:"accountNo"`
		NetworkType string `json:"network_type"`
		BinPrefix   string `json:"binPrefix"`
		AddressDtls struct {
			Address1 string `json:"address1"`
			Address2 string `json:"address2"`
			Address3 string `json:"address3"`
			City     string `json:"city"`
			State    string `json:"state"`
			Pincode  string `json:"pincode"`
			Country  string `json:"country"`
		} `json:"address_dtls"`
		SessionId string `json:"sessionId"`
	} `json:"msg"`
	Token  string `json:"token"`
	Signcs string `json:"signcs"`
}
