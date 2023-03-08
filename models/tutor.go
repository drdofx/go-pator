package models

import "gorm.io/gorm"

type Tutor struct {
	gorm.Model
	NumberOfTutees     int
	SelfDescription    string `gorm:"type:longtext;not null;"`
	AverageTutorRating float64
	UserID             uint `gorm:"unique, constraint:OnDelete:CASCADE"`
}
