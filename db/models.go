package db

import "time"

const TableNameUserDetails = "public.user_details"

// UserDetails mapped from table <public.user_details>
type UserDetails struct {
	ID                         int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt                  time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UserID                     string    `gorm:"column:user_id;not null;unique" json:"user_id"`
	Language                   string    `gorm:"column:language" json:"language"`
	Age                        int64     `gorm:"column:age" json:"age"`
	GrossMonthlyIncome         float64   `gorm:"column:gross_monthly_income" json:"gross_monthly_income"`
	EducationalQualification   string    `gorm:"column:educational_qualification" json:"educational_qualification"`
	Profession                 string    `gorm:"column:profession" json:"profession"`
	Gender                     string    `gorm:"column:gender" json:"gender"`
	MonthlyEssentialExpense    float64   `gorm:"column:monthly_essential_expense" json:"monthly_essential_expense"`
	MonthlyNonEssentialExpense float64   `gorm:"column:monthly_non_essential_expense" json:"monthly_non_essential_expense"`
	MonthlySavings             float64   `gorm:"column:monthly_savings" json:"monthly_savings"`
	MonthlyInvestments         float64   `gorm:"column:monthly_investments" json:"monthly_investments"`
	MonthlyInvestibleSurplus   float64   `gorm:"column:monthly_investible_surplus" json:"monthly_investible_surplus"`
	UpdatedAt                  time.Time `gorm:"column:updated_at;default:now()" json:"updated_at"`
}

// TableName UserDetails's table name
func (*UserDetails) TableName() string {
	return TableNameUserDetails
}
