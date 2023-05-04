package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sid-008/Postgres_CRUD/database"
	"github.com/sid-008/Postgres_CRUD/helper"
	"github.com/sid-008/Postgres_CRUD/models"
)

func AddPost(c *gin.Context) {
	var input models.Post
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.UserID = user.ID

	savedEntry, err := input.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}

func GetAllPosts(c *gin.Context) {
	user, err := helper.CurrentUser(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user.Posts})
}

func UpdateOnePost(c *gin.Context) {
	var update models.Post
	if err := c.ShouldBindJSON(&update); err != nil { // bind update request
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(c) //auth to get the current user

	if err != nil { //this is the auth, tested with random jwt, does not work if invalid signature
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post := user.Posts

	database.Database.Model(&post).Where("Title = ?", update.Title).Update("Content", update.Content) // this query updates based on the requested title
	//TODO add method that will also update Title
	c.JSON(http.StatusOK, gin.H{"updated": update})
}

func DeleteOnePost(c *gin.Context) {
	var deleted models.Post
	if err := c.ShouldBindJSON(&deleted); err != nil { // bind update request
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(c) //auth to get the current user

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post := user.Posts

	database.Database.Model(&post).Where("Title = ?", deleted.Title).Delete(&post) // this query deletes based on the requested title
	// This is a soft delete, adding .Unscoped method will make it a hard delete
	c.JSON(http.StatusOK, gin.H{"deleted": deleted})
}

func GetAllPostsAnon(c *gin.Context) {
	posts := models.GetAll()
	log.Println(posts) //TODO add error handling
	c.JSON(http.StatusOK, gin.H{"data": posts})
}
