package models

import "gorm.io/gorm"

type Word struct {
	gorm.Model

	Original string `gorm:"not null;unique_index" json:"original"`
	Local    string `gorm:"not null" json:"local"`
}
