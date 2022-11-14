package models

import "time"

type UserIntermValues struct {
	ReferenceNumber string    `gorm:"column:ref_no" json:"ref_no"`
	OtpRequiredId   string    `gorm:"column:otp_req_id" json:"otp_req_id"`
	UserID          string    `gorm:"column:user_id;not null" json:"user_id"`
	ID              int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt       time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
}

type UserCardCreateInformation struct {
	CardNetwork     string    `gorm:"column:card_network;not null" json:"card_network"`
	BinPrefix       string    `gorm:"column:bin_prefix;not null" json:"bin_prefix"`
	CardProdName    string    `gorm:"column:card_prod_name;not null" json:"card_prod_name"`
	CardProdCode    string    `gorm:"column:card_prod_code;not null" json:"card_prod_code"`
	AccountProdCode string    `gorm:"column:account_prod_code;not null" json:"account_prod_code"`
	AccountProdName string    `gorm:"column:account_prod_name;not null" json:"account_prod_name"`
	UserID          string    `gorm:"column:user_id;not null" json:"user_id"`
	ID              int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt       time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
}
