package models

import (
	"time"
)

const TableNameNeftTransfer = "neft_transfers"

// NeftTransfer mapped from table <neft_transfers>
type NeftTransfer struct {
	ID          int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	UserID      string    `gorm:"column:user_id" json:"user_id"`
	Remarks     string    `gorm:"column:remarks" json:"remarks"`
	AccountID   string    `gorm:"column:account_id" json:"account_id"`
	Amount      string    `gorm:"column:amount" json:"amount"`
	CreditorID  string    `gorm:"column:creditor_id" json:"creditor_id"`
	PaymentType string    `gorm:"column:payment_type" json:"payment_type"`
	NeftId      string    `gorm:"column:neft_id" json:"neft_id"`
}

// TableName NeftTransfer's table name
func (*NeftTransfer) TableName() string {
	return TableNameNeftTransfer
}
