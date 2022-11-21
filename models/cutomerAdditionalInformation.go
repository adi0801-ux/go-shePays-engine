package models

import "time"

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

type CustomerAdditionalInformation struct {
	PinCode           string    `gorm:"column:pin_code;not null" json:"pin_code"`
	DOB               string    `gorm:"column:dob;not null" json:"dob"`
	Salutation        string    `gorm:"column:salutation;not null" json:"salutation"`
	Country           string    `gorm:"column:country;not null" json:"country"`
	State             string    `gorm:"column:state;not null" json:"state"`
	City              string    `gorm:"column:city;not null" json:"city"`
	AddressLine1      string    `gorm:"column:address_line_1;not null" json:"address_line_1"`
	AddressLine2      string    `gorm:"column:address_line_2" json:"address_line_2"`
	MartialStatus     string    `gorm:"column:martial_status;not null" json:"martial_status"`
	IncomeCategory    string    `gorm:"column:income_category;not null" json:"income_category"`
	ProfessionalCode  string    `gorm:"column:professional_code;not null" json:"professional_code"`
	Relation          string    `gorm:"column:relation;not null" json:"relation"`
	CustomerEducation string    `gorm:"column:customer_education;not null" json:"customer_education"`
	FathersName       string    `gorm:"column:father_name;not null" json:"father_name"`
	MothersName       string    `gorm:"column:mother_name;not null" json:"mother_name"`
	UserId            string    `gorm:"column:user_id;not null" json:"user_id"`
	ID                int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt         time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
}
