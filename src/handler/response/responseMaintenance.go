package response

import (
	"github.com/gin-gonic/gin"
	"github.com/victorramos887/go_pdf/src/schemas"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message":  msg,
		"erorCode": code,
	})

}

func sendSuccess(ctx *gin.Context, code int, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message": "success",
		"data":    data,
	})
}

// SendError is an exported version of sendError
func SendError(ctx *gin.Context, code int, msg string) {
	sendError(ctx, code, msg)
}

// SendSuccess is an exported version of sendSuccess
func SendSuccess(ctx *gin.Context, code int, data interface{}) {
	sendSuccess(ctx, code, data)
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateMaintenanceRequest struct {
	Message string                      `json:"message"`
	Data    schemas.MaintenanceResponse `json:"data"`
}
