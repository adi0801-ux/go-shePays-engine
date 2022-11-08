package models

type AccountCreationAPI struct {
	Channelid  string  `json:"channelid"`
	Appdtls    AppDtls `json:"appdtls"`
	Devicedtls struct {
		Deviceid string `json:"deviceid"`
	} `json:"devicedtls"`
	Deviceidentifier DeviceIdentifier `json:"deviceidentifier"`
	Custcreddtls     struct {
		Userid   string `json:"userid"`
		Usertype string `json:"usertype"`
		Role     string `json:"role"`
		Cif      string `json:"cif"`
	} `json:"custcreddtls"`
	Accountcreationdtl struct {
		AccountNoString       string `json:"accountNoString"`
		AccountTitle          string `json:"accountTitle"`
		AcctCurrencyString    string `json:"acctCurrencyString"`
		BranchCode            string `json:"branchCode"`
		CustomerIDString      string `json:"customerIDString"`
		FlgJointHolderString  string `json:"flgJointHolderString"`
		FlgRestrictAcctString string `json:"flgRestrictAcctString"`
		FlgSCWaiveString      string `json:"flgSCWaiveString"`
		FlgTransactionType    string `json:"flgTransactionType"`
		MinorAcctStatusString string `json:"minorAcctStatusString"`
		ProductCodeString     string `json:"productCodeString"`
	} `json:"accountcreationdtl"`
	Token  string `json:"token"`
	Signcs string `json:"signcs"`
}

type AccountCreationAPIResponse struct {
	AccDtls struct {
		Accountno string `json:"accountno"`
	} `json:"acc_dtls"`
	Response string `json:"response"`
	Respcode string `json:"respcode"`
}
