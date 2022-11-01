package models

import (
	"time"
)

const TableNameSavingsAccount = "savings_account"

// SavingsAccount mapped from table <savings_account>
type SavingsAccount struct {
	ID                int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt         time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	UserID            string    `gorm:"column:user_id" json:"user_id"`
	AccountId         string    `gorm:"column:account_id" json:"account_id"`
	Ifsc              string    `gorm:"column:ifsc" json:"ifsc"`
	BankAccountNumber string    `gorm:"column:bank_account_number" json:"bank_account_number"`
}

// TableName SavingsAccount's table name
func (*SavingsAccount) TableName() string {
	return TableNameSavingsAccount
}
