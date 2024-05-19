package main

import (
	"backend-api/controllers"
	"backend-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})

	router.GET("/api/posts", controllers.FindPost)
	router.POST("/api/posts", controllers.StorePost)
	router.GET("/api/posts/:id", controllers.FindPostById)

	router.Run(":3000")
}
