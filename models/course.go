package models

import (
	"pator-api/database"

	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	CourseName  string `gorm:"type:varchar(255);not null;unique"`
	CourseProdi string `gorm:"type:varchar(255);not null;"`
}

func (c *Course) SaveCourse() (*Course, error) {
	err := database.DB.Create(&c).Error

	if err != nil {
		return &Course{}, err
	}

	return c, nil
}

func (c *Course) FindAllCourses() (*[]Course, error) {
	err := database.DB.Find(&c).Error

	if err != nil {
		return &[]Course{}, err
	}

	return &[]Course{}, nil
}

func (c *Course) FindCourseByID(id uint) (*Course, error) {
	err := database.DB.First(&c, id).Error

	if err != nil {
		return &Course{}, err
	}

	return c, nil
}

func (c *Course) UpdateCourse(id uint) (*Course, error) {
	err := database.DB.Model(&c).Where("id = ?", id).Updates(&c).Error

	if err != nil {
		return &Course{}, err
	}

	return c, nil
}
