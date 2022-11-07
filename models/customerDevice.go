package models

import "time"

type CustomerInformation struct {
	CustomerNickname  string `json:"customer_nickname"`
	CustomerName      string `json:"customer_name"`
	CustomerFName     string `json:"customer_f_name"`
	CustomerMName     string `json:"customer_m_name"`
	CustomerLName     string `json:"customer_l_name"`
	EmailId           string `json:"email_id"`
	Msisdn            string `json:"msisdn"`
	MobileCountryCode string `json:"mobile_country_code"`
	Gender            string `json:"gender"`
	UserId            string `json:"user_id"`
}

type CustomerDeviceApi struct {
	Mbappdetails struct {
		Channelid  string  `json:"channelid"`
		Appdtls    AppDtls `json:"appdtls"`
		Devicedtls struct {
			Deviceid string `json:"deviceid"`
		} `json:"devicedtls"`
		Deviceidentifier DeviceIdentifier `json:"deviceidentifier"`
		Custprofiledtls  struct {
			CustNickname      string `json:"cust_nickname"`
			CustName          string `json:"cust_name"`
			CustNameF         string `json:"cust_name_f"`
			CustNameM         string `json:"cust_name_m"`
			CustNameL         string `json:"cust_name_l"`
			Emailid           string `json:"emailid"`
			Emailidverflag    string `json:"emailidverflag"`
			Msisdn            string `json:"msisdn"`
			Mobilecountrycode string `json:"mobilecountrycode"`
			Mobileverflag     string `json:"mobileverflag"`
			Gender            string `json:"gender"`
			Mobverunqid       string `json:"mobverunqid"`
			Emailverunqid     string `json:"emailverunqid"`
		} `json:"custprofiledtls"`
	} `json:"mbappdetails"`
	Token  string `json:"token"`
	Signcs string `json:"signcs"`
}

type CustomerDeviceApiResponse struct {
	Deviceidentifier struct {
		Custunqid string `json:"custunqid"`
	} `json:"deviceidentifier"`
	Response string `json:"response"`
	Respcode string `json:"respcode"`
}

type CustomerDetails struct {
	CustomerNickname  string    `gorm:"column:customer_nickname" json:"customer_nickname"`
	CustomerName      string    `gorm:"column:customer_name;not null" json:"customer_name"`
	CustomerFName     string    `gorm:"column:customer_f_name" json:"customer_f_name"`
	CustomerMName     string    `gorm:"column:customer_m_name" json:"customer_m_name"`
	CustomerLName     string    `gorm:"column:customer_l_name" json:"customer_l_name"`
	EmailId           string    `gorm:"column:email_id;not null" json:"email_id"`
	Msisdn            string    `gorm:"column:msisdn;not null" json:"msisdn"`
	MobileCountryCode string    `gorm:"column:mobile_country_code;not null" json:"mobile_country_code"`
	Gender            string    `gorm:"column:gender" json:"gender"`
	UserID            string    `gorm:"column:user_id;not null" json:"user_id"`
	ID                int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt         time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
}

type CustomerRecordMobileNumberCheck struct {
	Channelid  string  `json:"channelid"`
	Appdtls    AppDtls `json:"appdtls"`
	Devicedtls struct {
		Deviceid string `json:"deviceid"`
	} `json:"devicedtls"`
	Deviceidentifier DeviceIdentifier `json:"deviceidentifier"`
	Searchtype       string           `json:"searchtype"`
	Searchstr        string           `json:"searchstr"`
	Token            string           `json:"token"`
	Signcs           string           `json:"signcs"`
}

type CustomerRecordMobileNumberCheckResponse struct {
	Response string `json:"response"`
	Respcode string `json:"respcode"`
	UserDtls struct {
		Cif                interface{} `json:"cif"`
		NbfcFlg            string      `json:"nbfc_flg"`
		Gender             string      `json:"gender"`
		Aadhaarno          string      `json:"aadhaarno"`
		Panno              interface{} `json:"panno"`
		Mobileno           string      `json:"mobileno"`
		AccountStatus      interface{} `json:"account_status"`
		Userid             string      `json:"userid"`
		CodAcctStat        interface{} `json:"cod_acct_stat"`
		CustrecId          int         `json:"custrec_id"`
		CustNameM          string      `json:"cust_name_m"`
		CustNameL          string      `json:"cust_name_l"`
		SokinFlg           interface{} `json:"sokin_flg"`
		AcntProdName       interface{} `json:"acnt_prod_name"`
		BankId             string      `json:"bank_id"`
		EsignFlg           string      `json:"esign_flg"`
		CustName           string      `json:"cust_name"`
		Useridstatus       interface{} `json:"useridstatus"`
		KycStatus          string      `json:"kyc_status"`
		CustNameF          string      `json:"cust_name_f"`
		Emailid            string      `json:"emailid"`
		JiffyFlg           interface{} `json:"jiffy_flg"`
		Facematch          string      `json:"facematch"`
		Custstatus         string      `json:"custstatus"`
		IfscCode           string      `json:"ifsc_code"`
		Kyctype            string      `json:"kyctype"`
		KycCustStatus      string      `json:"kyc_cust_status"`
		Dob                interface{} `json:"dob"`
		Accountno          interface{} `json:"accountno"`
		AcntProdCode       interface{} `json:"acnt_prod_code"`
		AllowedAccountCode interface{} `json:"allowed_account_code"`
		Salutation         string      `json:"salutation"`
		CustCatcode        string      `json:"cust_catcode"`
		Username           interface{} `json:"username"`
	} `json:"user_dtls"`
}

type CustomerAdditionalInformation struct {
	PinCode    string    `gorm:"column:pin_code" json:"pin_code"`
	DOB        string    `gorm:"column:dob" json:"dob"`
	Salutation string    `gorm:"column:salutation" json:"salutation"`
	Country    string    `gorm:"column:country" json:"country"`
	State      string    `gorm:"column:state" json:"state"`
	City       string    `gorm:"column:city" json:"city"`
	UserId     string    `gorm:"column:user_id" json:"user_id"`
	ID         int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt  time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
}
