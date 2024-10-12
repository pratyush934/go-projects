package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pratyush934/go-projects/curd-api/database"
	"github.com/pratyush934/go-projects/curd-api/models"
	"net/http"
)

type CreatingInput struct {
	Title   *string `json:"title" binding:"required"`
	Content *string `json:"content" binding:"required"`
}

func CreatePosts(c *gin.Context) {
	var input CreatingInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	post := models.Post{Title: input.Title, Content: input.Content}
	database.DB.Create(&post)

	c.JSON(http.StatusOK, gin.H{
		"message": "We did it successfully",
		"data":    post,
	})
}
