package models

type SetMPIN struct {
	CredVal     string `json:"cred_val"`
	CredSaltKey string `json:"cred_salt_key"`
	UserId      string `json:"user_id"`
}

type SetMPINAPI struct {
	Channelid  string  `json:"channelid"`
	Appdtls    AppDtls `json:"appdtls"`
	Devicedtls struct {
		Deviceid string `json:"deviceid"`
	} `json:"devicedtls"`
	Deviceidentifier DeviceIdentifier `json:"deviceidentifier"`
	Creddtls         struct {
		Credtype     string `json:"credtype"`
		Credval      string `json:"credval"`
		Credcategory string `json:"credcategory"`
		Credsaltkey  string `json:"credsaltkey"`
	} `json:"creddtls"`
	Custcreddtls struct {
		Userid   string `json:"userid"`
		Usertype string `json:"usertype"`
		Role     string `json:"role"`
	} `json:"custcreddtls"`
	Ivspec string `json:"ivspec"`
	Token  string `json:"token"`
	Signcs string `json:"signcs"`
}
