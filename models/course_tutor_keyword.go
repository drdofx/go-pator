package models

import "gorm.io/gorm"

type CourseTutorKeyword struct {
	gorm.Model
	Keyword       string `gorm:"type:varchar(255);not null;"`
	CourseTutorID uint   `gorm:"constraint:OnDelete:CASCADE"`
}
