package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pratyush934/go-projects/curd-api/controllers"
	"github.com/pratyush934/go-projects/curd-api/database"
	"log"
	"net/http"
)

const PORT string = ":8080"

func main() {

	router := gin.Default()

	database.ConnectDataBase()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusAccepted, gin.H{
			"message": "Hello I am Pratyush",
		})
	})
	router.GET("/get-post", controllers.GetPosts)
	router.POST("/post", controllers.CreatePosts)
	router.GET("/get-post/:id", controllers.GetPostById)
	router.PATCH("/post/:id", controllers.UpdatePost)
	router.DELETE("/post/:id", controllers.DeleteById)

	log.Fatal(router.Run(PORT))
}
