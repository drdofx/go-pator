package migrations

import (
	"pator-api/database"
	"pator-api/models"
)

func Migrate() {
	db := database.DB

	err := db.AutoMigrate(&models.User{}, &models.Tutor{}, &models.Tutee{})

	if err != nil {
		panic("Failed to migrate the database!")
	}

}
