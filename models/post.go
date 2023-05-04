package models

import (
	"log"

	"github.com/sid-008/Postgres_CRUD/database"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title   string `gorm:"type:text" json:"title"`
	Content string `gorm:"type:text" json:"content"`
	UserID  uint
}

func GetAll() []Post {
	var posts []Post
	err := database.Database.Select("content").Find(&posts).Error
	if err != nil {
		log.Fatal(err)
	}
	return posts
}
