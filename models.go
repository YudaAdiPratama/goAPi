package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"not null"`
	Password     string `gorm:"not null"`
	Email        string `gorm:"not null"`
	Name         string `gorm:"not null"`
	ProfileImage string // Assuming this is a URL or file path
}
