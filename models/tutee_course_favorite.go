package models

import (
	"time"

	"gorm.io/gorm"
)

type TuteeCourseFavorite struct {
	gorm.Model
	DateAdded     time.Time
	TuteeID       uint `gorm:"constraint:OnDelete:CASCADE"`
	CourseTutorID uint `gorm:"constraint:OnDelete:CASCADE"`
}
