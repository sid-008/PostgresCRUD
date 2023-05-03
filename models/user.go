package models

import (
	"html"
	"log"
	"strings"

	"github.com/sid-008/Postgres_CRUD/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null" json:"-"`
	Posts    []Post
}

func (user *User) Save() (*User, error) { //create user
	err := database.Database.Create(&user).Error
	if err != nil {
		log.Fatal(err)
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error { //password hashing method
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

func (user *User) ValidatePassword(password string) error { //pass validation
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByUsername(username string) (User, error) { //search by username, query db
	var user User
	err := database.Database.Where("username=?", username).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func FindUserById(id uint) (User, error) {
	var user User
	err := database.Database.Preload("Posts").Where("ID=?", id).Find(&user).Error // eager loading, this populates the entries slice in the user struct
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (post *Post) Save() (*Post, error) {
	err := database.Database.Create(&post).Error
	if err != nil {
		log.Fatal(err)
		return &Post{}, err
	}
	return post, nil
}
