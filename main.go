package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sid-008/Postgres_CRUD/controllers"
	"github.com/sid-008/Postgres_CRUD/database"
	"github.com/sid-008/Postgres_CRUD/models"
)

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(models.User{})
	database.Database.AutoMigrate(models.Post{})
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	loadEnv()
	loadDatabase()
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	router.GET("/register", controllers.Register)

	err := router.Run("localhost:3000")
	if err != nil {
		log.Fatal(err)
	}
}
