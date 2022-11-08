package models

type SMSGenrationAPI struct {
	Channelid  string  `json:"channelid"`
	Appdtls    AppDtls `json:"appdtls"`
	Devicedtls struct {
		Deviceid string `json:"deviceid"`
	} `json:"devicedtls"`
	Deviceidentifier DeviceIdentifier `json:"deviceidentifier"`
	Mobileno         string           `json:"mobileno"`
	Emailid          string           `json:"emailid"`
	Sentmode         string           `json:"sentmode"`
	Type             string           `json:"type"`
	Subtype          string           `json:"subtype"`
	Attempt          string           `json:"attempt"`
	Token            string           `json:"token"`
	Signcs           string           `json:"signcs"`
}

type UserId struct {
	UserId string `json:"user_id"`
}

type SMSGenrationAPIResponse struct {
	OtpDtls struct {
		OtprecId  float64     `json:"otprec_id"`
		Otpreqid  string      `json:"otpreqid"`
		Otpvalenc interface{} `json:"otpvalenc"`
		Otpexpdt  int64       `json:"otpexpdt"`
	} `json:"otp_dtls"`
	Response string `json:"response"`
	Respcode string `json:"respcode"`
}

type SMSVerifyAPI struct {
	Channelid  string  `json:"channelid"`
	Appdtls    AppDtls `json:"appdtls"`
	Devicedtls struct {
		Deviceid string `json:"deviceid"`
	} `json:"devicedtls"`
	Deviceidentifier DeviceIdentifier `json:"deviceidentifier"`
	Type             string           `json:"type"`
	Subtype          string           `json:"subtype"`
	Verstatusdtl     struct {
		Verunqid  string `json:"verunqid"`
		Smsactdtl struct {
			Smsto        string `json:"smsto"`
			Smsval       string `json:"smsval"`
			Smsactstatus string `json:"smsactstatus"`
		} `json:"smsactdtl"`
		Otpdtl struct {
			Otplength string `json:"otplength"`
			Otpval    string `json:"otpval"`
			Otpencval string `json:"otpencval"`
		} `json:"otpdtl"`
	} `json:"verstatusdtl"`
	Token  string `json:"token"`
	Signcs string `json:"signcs"`
}

type OTPVerify struct {
	UserId string `json:"user_id"`
	OTP    string `json:"otp"`
}

type SMSVerifyAPIResponse struct {
	Response string `json:"response"`
	Respcode string `json:"respcode"`
}
