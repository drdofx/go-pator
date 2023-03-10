package main

import (
	"pator-api/controllers/users"
	"pator-api/database"
	"pator-api/middlewares"
	"pator-api/migrations"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// load .env file
	err := godotenv.Load()

	if err != nil {
		panic("Failed to load .env file!")
	}

	database.Connect()
	migrations.Migrate()
}

func main() {
	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", users.Register)
	public.POST("/login", users.Login)

	protected := r.Group("/api/user", middlewares.JwtAuthMiddleware())
	protected.GET("/profile", users.Profile)

	r.Run()
}
