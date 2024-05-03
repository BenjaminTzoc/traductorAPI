package models

import "gorm.io/gorm"

type WordCategory struct {
	gorm.Model
	WordID     uint
	CategoryID uint
}
