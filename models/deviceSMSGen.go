package models

type SMSGenReq struct {
	Type        string `json:"type"`
	Subtype     string `json:"sub_type"`
	Attempt     int    `json:"attempt"`
	MeiNumber   string `json:"mei_number"`
	Imei        string `json:"imei"`
	SerialNo    string `json:"serial_no"`
	SimState    string `json:"sim_state"`
	SimOperator string `json:"sim_operator"`
	UserId      string `json:"user_id"`
}

type SMSGenReqAPI struct {
	Channelid  string  `json:"channelid"`
	Appdtls    AppDtls `json:"appdtls"`
	Devicedtls struct {
		Deviceid string `json:"deviceid"`
	} `json:"devicedtls"`
	Deviceidentifier DeviceIdentifier `json:"deviceidentifier"`
	Type             string           `json:"type"`
	Subtype          string           `json:"subtype"`
	Attempt          int              `json:"attempt"`
	Meinumber        string           `json:"meinumber"`
	Imei             string           `json:"imei"`
	Serialno         string           `json:"serialno"`
	Simstate         string           `json:"simstate"`
	Simoperator      string           `json:"simoperator"`
	Token            string           `json:"token"`
	Signcs           string           `json:"signcs"`
}

type SMSGenReqAPIResponse struct {
	VMOBREGSMSACTSMS   string `json:"V_MOBREG_SMSACT_SMS"`
	Response           string `json:"V_RESPONSE"`
	VMOBREGUNQID       string `json:"V_MOBREGUNQID"`
	VMOBREGSMSACTMOBTO string `json:"V_MOBREG_SMSACT_MOBTO"`
	Respcode           string `json:"V_RESPCODE"`
}

type SMSActionStatusAPI struct {
	Channelid  string  `json:"channelid"`
	Appdtls    AppDtls `json:"appdtls"`
	Devicedtls struct {
		Deviceid string `json:"deviceid"`
	} `json:"devicedtls"`
	Deviceidentifier DeviceIdentifier `json:"deviceidentifier"`
	Type             string           `json:"type"`
	Subtype          string           `json:"subtype"`
	Token            string           `json:"token"`
	Signcs           string           `json:"signcs"`
}

type SMSActionStatusAPIResponse struct {
	VSIMSTATE    string `json:"V_SIMSTATE"`
	Response     string `json:"V_RESPONSE"`
	VIMEI        string `json:"V_IMEI"`
	VMOBILENO    string `json:"V_MOBILENO"`
	VSTATUS      string `json:"V_STATUS"`
	Respcode     string `json:"V_RESPCODE"`
	VMEINUMBER   string `json:"V_MEINUMBER"`
	VSERIALNO    string `json:"V_SERIALNO"`
	VSIMOPERATOR string `json:"V_SIMOPERATOR"`
}
