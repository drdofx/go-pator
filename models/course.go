package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	CourseName  string `gorm:"type:varchar(255);not null;"`
	CourseProdi string `gorm:"type:varchar(255);not null;"`
}
