package router

import (
	"github.com/gin-gonic/gin"
	"github.com/victorramos887/go_pdf/src/handler"
)

func Initialize() {
	handler.InitializeHandler()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the API!",
		})
	})

	router.POST("/maintenance", handler.CreateMaintenance)
	router.Run(":8080")
}
