package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null;"`
	Email    string `gorm:"type:varchar(255);not null;unique"`
	Password string `gorm:"type:varchar(255);not null;" json:"-"`
	NIM      string `gorm:"type:varchar(15);not null;unique"`
	Major    string `gorm:"type:varchar(150);not null;"`
	Year     int    `gorm:"type:int;not null;"`
}
