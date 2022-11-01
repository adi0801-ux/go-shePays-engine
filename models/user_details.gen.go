package models

import (
	"time"
)

const TableNameUserDetail = "user_details"

// UserDetail mapped from table <user_details>
type UserDetail struct {
	ID               int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt        time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	FirstName        string    `gorm:"column:first_name;not null" json:"first_name"`
	LastName         string    `gorm:"column:last_name;not null" json:"last_name"`
	Gender           string    `gorm:"column:gender;not null" json:"gender"`
	Mobile           string    `gorm:"column:mobile;not null" json:"mobile"`
	Dob              string    `gorm:"column:dob;not null" json:"dob"`
	Email            string    `gorm:"column:email;not null" json:"email"`
	MiddleName       string    `gorm:"column:middle_name" json:"middle_name"`
	Title            string    `gorm:"column:title;not null" json:"title"`
	UpdatedAt        time.Time `gorm:"column:updated_at;default:now()" json:"updated_at"`
	UserID           string    `gorm:"column:user_id" json:"user_id"`
	EmploymentStatus string    `gorm:"column:employment_status" json:"employment_status"`
	MaritalStatus    string    `gorm:"column:marital_status" json:"marital_status"`
	Nationality      string    `gorm:"column:nationality" json:"nationality"`
	HappayUserId     string    `gorm:"column:happay_user_id" json:"happay_user_id"`
	KycStatus        string    `gorm:"column:kyc_status" json:"kyc_status"`
	MobileVerified   bool      `gorm:"column:mobile_verified" json:"mobile_verified"`
}

// TableName UserDetail's table name
func (*UserDetail) TableName() string {
	return TableNameUserDetail
}

//happay api

type UserDetailsApi struct {
	Gender     string `json:"gender"`
	LastName   string `json:"last_name"`
	Mobile     string `json:"mobile"`
	Dob        string `json:"dob"`
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	Title      string `json:"title"`
	AddressId  string `json:"address_id"`
}
