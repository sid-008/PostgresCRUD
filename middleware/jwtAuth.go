package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sid-008/Postgres_CRUD/helper"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := helper.ValidateJWT(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Auth required bro"})
			c.Abort()
			return
		}
		c.Next()
	}
}
