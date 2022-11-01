package models

import (
	"time"
)

const TableNameDispute = "disputes"

// Dispute mapped from table <disputes>
type Dispute struct {
	ID                    int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt             time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	UserID                string    `gorm:"column:user_id" json:"user_id"`
	Comment               string    `gorm:"column:comment" json:"comment"`
	DisputeAmount         string    `gorm:"column:dispute_amount" json:"dispute_amount"`
	DisputeAmountCurrency string    `gorm:"column:dispute_amount_currency" json:"dispute_amount_currency"`
	Reason                string    `gorm:"column:reason" json:"reason"`
	TransactionID         string    `gorm:"column:transaction_id" json:"transaction_id"`
}

// TableName Dispute's table name
func (*Dispute) TableName() string {
	return TableNameDispute
}
