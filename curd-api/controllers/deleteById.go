package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pratyush934/go-projects/curd-api/database"
	"github.com/pratyush934/go-projects/curd-api/models"
	"net/http"
)

func DeleteById(c *gin.Context) {

	var post models.Post

	id := c.Param("id")

	if err := database.DB.Where("id = ?", id).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"err": err,
		})
	}
	database.DB.Delete(&post)

	c.JSON(http.StatusOK, gin.H{
		"data-deleted": post,
	})
}
