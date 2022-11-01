package models

import (
	"time"
)

const TableNameUserCkyc = "user_ckyc"

// UserCkyc mapped from table <user_ckyc>
type UserCkyc struct {
	ID           int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt    time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	UserID       string    `gorm:"column:user_id" json:"user_id"`
	DocID        string    `gorm:"column:doc_id" json:"doc_id"`
	DocType      string    `gorm:"column:doc_type" json:"doc_type"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
	PlaceOfIssue string    `gorm:"column:place_of_issue" json:"place_of_issue"`
	IssueDate    string    `gorm:"column:issue_date" json:"issue_date"`
	CkycDate     string    `gorm:"column:ckyc_date" json:"ckyc_date"`
	CkycName     string    `gorm:"column:ckyc_name" json:"ckyc_name"`
	CkycNumber   string    `gorm:"column:ckyc_no" json:"ckyc_no"`
}

// TableName UserCkyc's table name
func (*UserCkyc) TableName() string {
	return TableNameUserCkyc
}
