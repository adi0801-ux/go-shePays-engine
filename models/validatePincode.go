package models

type ValidatePincode struct {
	Channelid  string  `json:"channelid"`
	Appdtls    AppDtls `json:"appdtls"`
	Devicedtls struct {
		Deviceid string `json:"deviceid"`
	} `json:"devicedtls"`
	Deviceidentifier DeviceIdentifier `json:"deviceidentifier"`
	Pincode          string           `json:"pincode"`
	Token            string           `json:"token"`
	Signcs           string           `json:"signcs"`
}

type ValidatePincodeResponse struct {
	Response    string `json:"response"`
	RespMessage string `json:"RespMessage"`
	Respcode    string `json:"respcode"`
	Citydetails []struct {
		Pincode           string `json:"pincode"`
		Relatedheadoffice string `json:"relatedheadoffice"`
		Taluk             string `json:"taluk"`
		Divisionname      string `json:"divisionname"`
		Latitude          string `json:"latitude"`
		Telephone         string `json:"telephone"`
		Statename         string `json:"statename"`
		Officetype        string `json:"officetype"`
		Circlename        string `json:"circlename"`
		Mfstatecode       string `json:"mfstatecode"`
		Officename        string `json:"officename"`
		Deliverystatus    string `json:"deliverystatus"`
		Districtname      string `json:"districtname"`
		Relatedsuboffice  string `json:"relatedsuboffice"`
		Regionname        string `json:"regionname"`
		Longitude         string `json:"longitude"`
		Mfcountrycode     string `json:"mfcountrycode"`
	} `json:"citydetails"`
}
