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

	err := database.Database.AutoMigrate(models.User{})
	if err != nil {
		log.Fatal(err)
	}

	err = database.Database.AutoMigrate(models.Post{})
	if err != nil {
		log.Fatal(err)
	}

}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal(err)
	}
}

func serveapp() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	err := router.Run("localhost:3000")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	loadEnv()
	loadDatabase()
	serveapp()
}
