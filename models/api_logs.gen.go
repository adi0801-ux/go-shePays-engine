package models

import (
	"time"
)

const TableNameAPILog = "api_logs"

// APILog mapped from table <api_logs>
type APILog struct {
	ID           int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	RequestId    string    `gorm:"column:request_id;not null" json:"request_id"`
	CreatedAt    time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	Params       string    `gorm:"column:params" json:"params"`
	Payload      string    `gorm:"column:payload" json:"payload"`
	Method       string    `gorm:"column:method" json:"method"`
	Endpoint     string    `gorm:"column:endpoint" json:"endpoint"`
	Response     string    `gorm:"column:response" json:"response"`
	ResponseCode int       `gorm:"column:response_code" json:"response_code"`
	Error        string    `gorm:"column:error" json:"error"`
	ResponseDate string    `gorm:"column:response_date" json:"response_date"`
}

// TableName APILog's table name
func (*APILog) TableName() string {
	return TableNameAPILog
}
