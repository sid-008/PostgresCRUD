package helper

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sid-008/Postgres_CRUD/models"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenJWT(user models.User) (string, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	return token.SignedString(privateKey)
}

func ValidateJWT(c *gin.Context) error {
	token, err := getToken(c)
	if err != nil {
		return err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token provided")
}

func getToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}

func getTokenFromRequest(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Cookie")
	log.Println(bearerToken)
	splitToken := strings.Split(bearerToken, "=")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}

func CurrentUser(c *gin.Context) (models.User, error) {
	err := ValidateJWT(c)
	if err != nil {
		return models.User{}, err
	}
	token, _ := getToken(c)
	claims, _ := token.Claims.(jwt.MapClaims)
	userId := uint(claims["id"].(float64))

	user, err := models.FindUserById(userId)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
