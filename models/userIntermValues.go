package models

import "time"

type UserIntermValues struct {
	ReferenceNumber string    `gorm:"column:ref_no" json:"ref_no"`
	OtpRequiredId   string    `gorm:"column:otp_req_id" json:"otp_req_id"`
	UserID          string    `gorm:"column:user_id;not null" json:"user_id"`
	ID              int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt       time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
}
