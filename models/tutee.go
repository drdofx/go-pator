package models

import "gorm.io/gorm"

type Tutee struct {
	gorm.Model
	AverageRatingGiven float64
	UserID             uint `gorm:"unique;constraint:OnDelete:CASCADE"`
}
