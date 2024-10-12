package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pratyush934/go-projects/curd-api/database"
	"github.com/pratyush934/go-projects/curd-api/models"
	"net/http"
)

func GetPostById(c *gin.Context) {
	var post models.Post

	id := c.Param("id")

	err := database.DB.Where("id = ?", id).First(&post).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": post,
	})
}
