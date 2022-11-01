package models

import (
	"time"
)

const TableNameUserCard = "user_cards"

// UserCard mapped from table <user_cards>
type UserCard struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	AddressID string    `gorm:"column:address_id" json:"address_id"`
	BinID     string    `gorm:"column:bin_id" json:"bin_id"`
	Status    string    `gorm:"column:status" json:"status"`
	Pin       string    `gorm:"column:pin" json:"pin"`
	IsLost    string    `gorm:"column:is_lost" json:"is_lost"`
	UserID    string    `gorm:"column:user_id;not null" json:"user_id"`
	CID       string    `gorm:"column:cid;not null" json:"cid"` // card id --> card_user_17783
}

// TableName UserCard's table name
func (*UserCard) TableName() string {
	return TableNameUserCard
}
