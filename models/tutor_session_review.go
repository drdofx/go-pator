package models

import "gorm.io/gorm"

type TutorSessionReview struct {
	gorm.Model
	TuteeReview    string `gorm:"type:longtext;not null;"`
	TuteeRating    int    `gorm:"type:int;not null;"`
	TutorComment   string `gorm:"type:longtext;not null;"`
	TutorSessionID uint   `gorm:"constraint:OnDelete:CASCADE"`
}
