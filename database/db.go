package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// load .env file
	err := godotenv.Load()

	if err != nil {
		panic("Failed to load .env file!")
	}

	// DB connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))

	db, errConnect := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if errConnect != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Database connection successfully opened")
}
