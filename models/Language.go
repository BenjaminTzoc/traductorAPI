package models

import "gorm.io/gorm"

type Language struct {
	gorm.Model
	Name    string `gorm:"not null"`
	Code    string `gorm:"not null"`
	IsLocal bool
}
