package models

type Company struct {
	ID    uint `gorm:"primarykey"`
	Name  string
	Users []User `gorm:"constraint:OnDelete:CASCADE;"`
}
