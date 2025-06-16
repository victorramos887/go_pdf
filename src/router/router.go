package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the API!",
		})
	})
	router.Run(":8080") // Start the server on port 8080
}
