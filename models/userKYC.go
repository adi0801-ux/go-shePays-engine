package models

import (
	"mime/multipart"
	"time"
)

type KYCPAN struct {
	PanNumber string    `gorm:"column:pan_number;not null" json:"pan_number"`
	UserID    string    `gorm:"column:user_id;not null" json:"user_id"`
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
}

type PANCheckExitsAPI struct {
	Channelid  string  `json:"channelid"`
	Appdtls    AppDtls `json:"appdtls"`
	Devicedtls struct {
		Deviceid string `json:"deviceid"`
	} `json:"devicedtls"`
	Deviceidentifier DeviceIdentifier `json:"deviceidentifier"`
	Panno            string           `json:"panno"`
	Token            string           `json:"token"`
	Signcs           string           `json:"signcs"`
}

type PANVerifyAPIResponse struct {
	Custdocresppoa struct {
		Custdocrecid float64 `json:"Custdocrecid"`
		Response     string  `json:"response"`
		Respcode     string  `json:"respcode"`
	} `json:"custdocresppoa"`
	Response   string `json:"response"`
	Respcode   string `json:"respcode"`
	Pandetails struct {
		Fname       string `json:"fname"`
		Lname       string `json:"lname"`
		DocXml      string `json:"doc_xml"`
		Mname       string `json:"mname"`
		Fullname    string `json:"fullname"`
		Salutation  string `json:"salutation"`
		DocId       string `json:"doc_id"`
		Issuedate   string `json:"issuedate"`
		Status      string `json:"status"`
		Adharseeded string `json:"adharseeded"`
	} `json:"pandetails"`
}

type KYCAadharVerify struct {
	AddharFile *multipart.FileHeader `json:"addhar_file"`
	Password   string                `json:"password"`
	UserId     string                `json:"user_id"`
}

type KYCAadharAPIResponse struct {
	Custdocresppoa struct {
		Custdocrecid float64 `json:"Custdocrecid"`
		Response     string  `json:"response"`
		Respcode     string  `json:"respcode"`
	} `json:"custdocresppoa"`
	AadhaarRefId string      `json:"aadhaar_refId"`
	Aadhaarno    string      `json:"aadhaarno"`
	Signature    string      `json:"signature"`
	Purpose      interface{} `json:"purpose"`
	Aadhaarrefno string      `json:"aadhaarrefno"`
	Respcode     string      `json:"respcode"`
	Uidrefid     string      `json:"uidrefid"`
	Custrecid    string      `json:"custrecid"`
	Doctype      string      `json:"doctype"`
	Response     string      `json:"response"`
	Docdtl1      string      `json:"docdtl1"`
	Aadhardtls   struct {
		Uid string `json:"uid"`
		Poa struct {
			Co      string `json:"co"`
			Street  string `json:"street"`
			House   string `json:"house"`
			Lm      string `json:"lm"`
			Loc     string `json:"loc"`
			Vtc     string `json:"vtc"`
			Subdist string `json:"subdist"`
			Dist    string `json:"dist"`
			State   string `json:"state"`
			Pc      string `json:"pc"`
			Po      string `json:"po"`
			Country string `json:"country"`
			VtcCode string `json:"vtcCode"`
		} `json:"poa"`
		LData interface{} `json:"lData"`
		Ldata string      `json:"ldata"`
		Poi   struct {
			Name   string `json:"name"`
			Dob    string `json:"dob"`
			Gender string `json:"gender"`
			Phone  string `json:"phone"`
			Email  string `json:"email"`
		} `json:"poi"`
		Pht string `json:"pht"`
		Prn string `json:"prn"`
	} `json:"aadhardtls"`
	Custdocresppoi struct {
		Custdocrecid float64 `json:"Custdocrecid"`
		Response     string  `json:"response"`
		Respcode     string  `json:"respcode"`
	} `json:"custdocresppoi"`
	Channelid         string `json:"channelid"`
	Aadhaarvaultrecid int    `json:"aadhaarvaultrecid"`
}

type ValidatePANAPI struct {
	Channelid  string  `json:"channelid"`
	Appdtls    AppDtls `json:"appdtls"`
	Devicedtls struct {
		Deviceid string `json:"deviceid"`
	} `json:"devicedtls"`
	Deviceidentifier DeviceIdentifier `json:"deviceidentifier"`
	Panno            string           `json:"panno"`
	Token            string           `json:"token"`
	Signcs           string           `json:"signcs"`
}

type ValidatePANAPIResponse struct {
	Response   string `json:"response"`
	Pandetails struct {
		Fname      string `json:"fname"`
		Lname      string `json:"lname"`
		DocXml     string `json:"doc_xml"`
		Mname      string `json:"mname"`
		Fullname   string `json:"fullname"`
		Salutation string `json:"salutation"`
		DocId      string `json:"doc_id"`
		Issuedate  string `json:"issuedate"`
		Status     string `json:"status"`
	} `json:"pandetails"`
	Respcode string `json:"respcode"`
}
