package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	log.Fatal(router.Run("localhost:8080"))
}
