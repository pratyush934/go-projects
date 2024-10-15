package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pratyush934/go-projects/curd-api/database"
	"github.com/pratyush934/go-projects/curd-api/models"
	"net/http"
)

type UpdatedPost struct {
	Title   *string `json:"title" binding:"required"`
	Content *string `json:"content" binding:"required"`
}

func UpdatePost(c *gin.Context) {

	var post models.Post
	//good going
	id := c.Param("id")

	if err := database.DB.Where("id = ?", id).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var input UpdatedPost

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	updatedInput := models.Post{Title: input.Title, Content: input.Content}

	database.DB.Model(&post).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{
		"updated-data": updatedInput,
	})
}
