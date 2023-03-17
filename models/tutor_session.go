package models

import (
	"time"

	"gorm.io/gorm"
)

type TutorSession struct {
	gorm.Model
	DateStarted   time.Time
	DateEnded     time.Time
	TuteeID       uint `gorm:"constraint:OnDelete:CASCADE"`
	TutorID       uint `gorm:"constraint:OnDelete:CASCADE"`
	CourseTutorID uint `gorm:"constraint:OnDelete:CASCADE"`
}
