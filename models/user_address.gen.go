package models

import (
	"time"
)

const TableNameUserAddress = "user_address"

// UserAddress mapped from table <user_address>
type UserAddress struct {
	ID           int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt    time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	AddressLine1 string    `gorm:"column:address_line_1;not null" json:"address_line_1"`
	AddressLine2 string    `gorm:"column:address_line_2" json:"address_line_2"`
	City         string    `gorm:"column:city;not null" json:"city"`
	Country      string    `gorm:"column:country;not null" json:"country"`
	State        string    `gorm:"column:state;not null" json:"state"`
	ZipCode      string    `gorm:"column:state;not null" json:"zip_code"`
	UserID       string    `gorm:"column:user_id;not null" json:"user_id"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
	AddressID    string    `gorm:"column:address_id" json:"address_id"`
}

// TableName UserAddress's table name
func (*UserAddress) TableName() string {
	return TableNameUserAddress
}

type UserAddressApi struct {
	AddressLine2 string       `json:"address_line_2"`
	City         string       `json:"city"`
	Country      string       `json:"country"`
	State        string       `json:"state"`
	ZipCode      string       `json:"zip_code"`
	AddressLine1 string       `json:"address_line_1"`
	ShippingInfo ShippingInfo `json:"shipping_info"`
}

type ShippingInfo struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Mobile    string `json:"mobile"`
}

type UserAddressInput struct {
	AddressLine2 string `json:"address_line_2"`
	City         string `json:"city"`
	Country      string `json:"country"`
	State        string `json:"state"`
	ZipCode      string `json:"zip_code"`
	AddressLine1 string `json:"address_line_1"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Mobile       string `json:"mobile"`
	UserID       string `json:"user_id"`
}
