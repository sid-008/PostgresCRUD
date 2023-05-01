package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	err := router.Run("localhost:3000")
	if err != nil {
		log.Fatal(err)
	}
}
