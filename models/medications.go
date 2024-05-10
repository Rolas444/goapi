package models

type Medication struct {
	ID      int64  `gorm:"column:id_med"`
	Name    string `gorm:"column:name"`
	Details string `gorm:"column:details"`
}

func (Medication) TableName() string {
	return "medications"
}
