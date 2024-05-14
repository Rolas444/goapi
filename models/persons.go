package models

type Person struct {
	// gorm.Model
	ID         int64  `gorm:"column:id_person;primary_key"`
	DocNumber  string `gorm:"column:doc_number"`
	Name       string `gorm:"column:name"`
	MiddleName string `gorm:"column:middle_name"`
	LastName   string `gorm:"column:last_name"`
}

func (Person) TableName() string {
	return "persons"
}
