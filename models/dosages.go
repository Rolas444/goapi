package models

type Dosage struct {
	ID             int64  `gorm:"column:id_dosage"`
	PrescriptionID int64  `gorm:"column:prescription_id"`
	MedicationID   int64  `gorm:"column:medication_id"`
	Detail         string `gorm:"column:detail"`
}

func (Dosage) TableName() string {
	return "dosages"
}
