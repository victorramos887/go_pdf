package router

import (
	"github.com/gin-gonic/gin"
	"github.com/victorramos887/go_pdf/src/handler"
)

func Initialize() {
	handler.InitializeHandler()

	router := gin.Default()
	initializeRoutes(router)
	router.Run(":8080")
}
