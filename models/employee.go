package models

import "time"

type Employee struct {
	FirstName       string    `gorm:"type:varchar(10)" json:"FirstName"`
	LastName        string    `gorm:"type:varchar(10)" json:"LastName"`
	ID              int       `gorm:"primarykey" json:"ID"`
	HireDate        time.Time `gorm:"type:date" json:"HireDate"`
	TerminationDate time.Time `gorm:"type:date" json:"TerminationDate"`
	Salary          int       `gorm:"type:int" json:"Salary"`
}
