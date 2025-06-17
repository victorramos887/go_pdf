package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/victorramos887/go_pdf/src/handler/response"
	"github.com/victorramos887/go_pdf/src/schemas"
)

func ListMaintenance(ctx *gin.Context) {
	maintenances := []schemas.Maintenance{}

	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)

	if err != nil || limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	if err := db.Limit(limit).Offset(offset).Find(&maintenances).Error; err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "error listing maintenance records")
		return
	}

	var total int64
	db.Model(&schemas.Maintenance{}).Count(&total)
	paginatedResponse := gin.H{
		"data":       maintenances,
		"page":       page,
		"limit":      limit,
		"total":      total,
		"totalPages": int((total + int64(limit) - 1) / int64(limit)),
	}
	response.SendSuccess(ctx, http.StatusOK, paginatedResponse)
}
