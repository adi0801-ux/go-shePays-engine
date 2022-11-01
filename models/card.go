package models

import "time"

type CreatePhysicalCard struct {
	BinId         string `json:"bin_id"`
	EmbossingName string `json:"embossing_name"`
	UserId        string `json:"user_id"`
}

type CreatePhysicalCardApi struct {
	BinId             string   `json:"bin_id"`
	EmbossingNameList []string `json:"embossing_name_list"`
	NumberOfCards     int      `json:"number_of_cards"`
	AddressId         string   `json:"address_id"`
}

const TableNamePhysicalCard = "physical_card"

type CreatePhysicalCardApiResponse struct {
	Status         string    `gorm:"column:status;not null" json:"status"`
	EmbossingName4 string    `gorm:"column:embossing_name4;not null" json:"embossing_name4"`
	EmbossingName3 string    `gorm:"column:embossing_name3;not null" json:"embossing_name3"`
	CardId         string    `gorm:"column:card_id;not null" json:"card_id"`
	CardFileId     string    `gorm:"column:card_file_id;not null" json:"card_file_id"`
	KitId          string    `gorm:"column:kit_id;not null" json:"kit_id"`
	Last4Digits    string    `gorm:"column:last_4_digits;not null" json:"last_4_digits"`
	Active         bool      `gorm:"column:active;not null" json:"active"`
	UserID         string    `gorm:"column:user_id;not null" json:"user_id"`
	ID             int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt      time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
}

func (*CreatePhysicalCardApiResponse) TableName() string {
	return TableNamePhysicalCard
}

type CreateVirtualCard struct {
	BinId  string `json:"bin_id"`
	UserId string `json:"user_id"`
}

const TableNameVirtualCard = "virtual_card"

type CreateVirtualCardApiResponse struct {
	Status         string    `gorm:"column:status;not null" json:"status"`
	EmbossingName4 string    `gorm:"column:embossing_name4;not null" json:"embossing_name4"`
	CardId         string    `gorm:"column:card_id;not null" json:"card_id"`
	EmbossingName3 string    `gorm:"column:embossing_name3;not null" json:"embossing_name3"`
	KitId          string    `gorm:"column:kit_id;not null" json:"kit_id"`
	Last4Digits    string    `gorm:"column:last_4_digits;not null" json:"last_4_digits"`
	Active         bool      `gorm:"column:active;not null" json:"active"`
	UserID         string    `gorm:"column:user_id;not null" json:"user_id"`
	ID             int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt      time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
}

func (*CreateVirtualCardApiResponse) TableName() string {
	return TableNameVirtualCard
}

type CreateVirtualCardApi struct {
	BinId string `json:"bin_id"`
}

type AssignCard struct {
	UserId string `json:"user_id"`
}

type InitiateCardPinSet struct {
	CardToken string `json:"card_token"`
}

type InitiateCardPinSetApiResponse struct {
	RedirectUrl string `json:"redirect_url"`
}

type SetCardPin struct {
	Pin    string `json:"pin"`
	UserId string `json:"user_id"`
}

type SetCardPinApi struct {
	Pin         string `json:"pin"`
	PinSetToken string `json:"pin_set_token"`
}

type ReplaceCard struct {
	NewCardId string `json:"new_card_id"`
	OldCardId string `json:"old_card_id"`
	UserId    string `json:"user_id"`
}
