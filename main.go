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
	authRoutes := router.Group("/auth")
	protectedRoutes := router.Group("/api")
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	router.GET("/all", controllers.GetAllPostsAnon) //anon readers can access through this

	authRoutes.POST("/register", controllers.Register)
	authRoutes.POST("/login", controllers.Login)
	authRoutes.POST("/logout", controllers.Logout) //just delete cookie with the jwt

	protectedRoutes.POST("/post", controllers.AddPost)
	protectedRoutes.GET("/post", controllers.GetAllPosts) //get all posts for currently logged in user
	protectedRoutes.PUT("/post", controllers.UpdateOnePost)
	protectedRoutes.DELETE("/post", controllers.DeleteOnePost)

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
