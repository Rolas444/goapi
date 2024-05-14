package models

type Prescription struct {
	// gorm.Model
	ID       int64  `gorm:"column:id_prescription"`
	PersonID int64  `gorm:"column:id_person"`
	Person   Person `gorm:"foreignKey:PersonID"`
}

func (Prescription) TableName() string {
	return "prescriptions"
}
