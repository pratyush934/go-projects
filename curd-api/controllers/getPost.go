package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pratyush934/go-projects/curd-api/database"
	"github.com/pratyush934/go-projects/curd-api/models"
	"net/http"
)

func GetPosts(c *gin.Context) {
	var posts []models.Post
	database.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{
		"data": posts,
	})
}
