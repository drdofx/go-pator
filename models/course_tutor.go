package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type CourseTutor struct {
	gorm.Model
	HourlyRate        decimal.Decimal `gorm:"type:decimal(10,2);not null;"`
	NumberOfTutees    int
	CourseDescription string `gorm:"type:longtext;not null;"`
	CourseRating      float64
	CourseID          uint `gorm:"constraint:OnDelete:CASCADE"`
	TutorID           uint `gorm:"constraint:OnDelete:CASCADE"`
}
