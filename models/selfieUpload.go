package models

type Selfie struct {
	Image   string `json:"image"`
	UserId  string `json:"user_id"`
	ReferId string `json:"refer_id"`
	JobId   string `json:"job_id"`
}

type VerifySelfieAPI struct {
	Channelid  string  `json:"channelid"`
	Appdtls    AppDtls `json:"appdtls"`
	Devicedtls struct {
		Deviceid string `json:"deviceid"`
	} `json:"devicedtls"`
	Deviceidentifier DeviceIdentifier `json:"deviceidentifier"`
	SelfieImage      string           `json:"selfieImage"`
	Referid          string           `json:"referid"`
	Jobid            string           `json:"jobid"`
	Token            string           `json:"token"`
	Signcs           string           `json:"signcs"`
}

type VerifySelfieAPIResponse struct {
	Response string `json:"response"`
	Facedtl  struct {
		Jobid     string `json:"jobid"`
		Referid   string `json:"referid"`
		Facematch string `json:"facematch"`
	} `json:"facedtl"`
	Respcode string `json:"respcode"`
}
