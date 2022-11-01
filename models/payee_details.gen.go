package models

import (
	"time"
)

const TableNamePayeeDetail = "payee_details"

// PayeeDetail mapped from table <payee_details>
type PayeeDetail struct {
	ID            int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt     time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	UserID        string    `gorm:"column:user_id" json:"user_id"`
	AccountNumber string    `gorm:"column:account_number" json:"account_number"`
	Ifsc          string    `gorm:"column:ifsc" json:"ifsc"`
	Name          string    `gorm:"column:name" json:"name"`
}

// TableName PayeeDetail's table name
func (*PayeeDetail) TableName() string {
	return TableNamePayeeDetail
}
