package router

import (
	"github.com/gin-gonic/gin"
	"github.com/victorramos887/go_pdf/src/handler"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	basePath := "/api/v1"
	v1 := router.Group(basePath)

	v1.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the API!",
		})
	})

	v1.POST("/maintenance", handler.CreateMaintenance)

	v1.GET("/maintenance", handler.ShowMaintenance)
}
