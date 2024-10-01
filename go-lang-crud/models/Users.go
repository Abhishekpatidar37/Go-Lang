package models

import (
	"golang-crud/enum"
	"time"
)

// User has many posts
type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	Name      string    `gorm:"size:100;not null"`
	Email     string    `gorm:"size:100;unique;not null" validate:"required,email"`
	Password  string    `gorm:"not null" validate:"required,min=8"`
	Role      enum.Role `gorm:"type:user_role;default:'user'"`
	CompanyID uint
	Company   Company
	Posts     []Post `gorm:"constraint:OnDelete:CASCADE;"`
}
