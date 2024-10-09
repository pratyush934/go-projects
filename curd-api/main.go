package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const PORT string = ":8080"

func main() {

	router := gin.Default()

	router.GET("/ping", func(g *gin.Context) {
		g.JSON(http.StatusAccepted, gin.H{
			"message": "pong",
		})
	})

	log.Fatal(router.Run(PORT))
}
