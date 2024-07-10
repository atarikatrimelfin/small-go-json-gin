package models

import "time"

type AnnualReviews struct {
	ID         int       `gorm:"primarykey" json:"ID"`
	EmpID      int       `gorm:"type:int" json:"empID"`
	ReviewDate time.Time `gorm:"type:date" json:"ReviewDate"`
}
