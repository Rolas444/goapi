package models

import "time"

type Prescription struct {
	ID           int64     `gorm:"column:id_prescription"`
	PersonID     int64     `gorm:"column:person_id"`
	CreationDate time.Time `gorm:"column:creation_date"`
}

func (Prescription) TableName() string {
	return "prescriptions"
}
