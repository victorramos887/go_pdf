package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victorramos887/go_pdf/src/handler/response"
	"github.com/victorramos887/go_pdf/src/schemas"
)

func ShowMaintenance(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		response.SendError(ctx, http.StatusBadRequest, "id parameter is required")
		return
	}

	maintenance := schemas.Maintenance{}

	if err := db.First(&maintenance, id).Error; err != nil {
		response.SendError(ctx, http.StatusNotFound, fmt.Sprintf("maintenance with id %s not found", id))
		return
	}

	response.SendSuccess(ctx, http.StatusOK, maintenance)

}
