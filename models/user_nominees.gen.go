package models

import (
	"time"
)

const TableNameUserNominee = "user_nominees"

// UserNominee mapped from table <user_nominees>
type UserNominee struct {
	ID           int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt    time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	NomineeName  string    `gorm:"column:nominee_name" json:"nominee_name"`
	RelationType string    `gorm:"column:relation_type" json:"relation_type"`
	UserID       string    `gorm:"column:user_id" json:"user_id"`
	AddressLine1 string    `gorm:"column:address_line_1" json:"address_line_1"`
	City         string    `gorm:"column:city" json:"city"`
	Country      string    `gorm:"column:country" json:"country"`
	PostalCode   string    `gorm:"column:postal_code" json:"postal_code"`
	State        string    `gorm:"column:state" json:"state"`
}

// TableName UserNominee's table name
func (*UserNominee) TableName() string {
	return TableNameUserNominee
}

type UserNomineesInput struct {
	Address struct {
		AddressLine1 string `json:"address_line_1"`
		City         string `json:"city"`
		Country      string `json:"country"`
		PostalCode   string `json:"postal_code"`
		State        string `json:"state"`
	} `json:"address"`
	NomineeName  string `json:"nominee_name"`
	RelationType string `json:"relation_type"`
	Value        int    `json:"value"`
}

type UserNomineesApi struct {
	UserNominees []UserNomineesInput `json:"user_nominees"`
	UserId       string              `json:"user_id"`
}
