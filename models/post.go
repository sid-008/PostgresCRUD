package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Content string `gorm:"type:text" json:"content"`
	UserID  uint
}
