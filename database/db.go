package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// DB connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	db, errConnect := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if errConnect != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Database connection successfully opened")

	DB = db
}
