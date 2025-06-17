package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victorramos887/go_pdf/src/handler/request"
	"github.com/victorramos887/go_pdf/src/handler/response"
	"github.com/victorramos887/go_pdf/src/schemas"
)

func CreateMaintenance(ctx *gin.Context) {
	requestBody := request.CreateMaintenanceRequest{}

	ctx.BindJSON(&requestBody)

	if err := requestBody.Validate(); err != nil {
		response.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	maintenance := schemas.Maintenance{
		Name: requestBody.Name,
		Data: requestBody.Data,
		Hora: requestBody.Hora,
	}

	fmt.Println("Instancia do banco de dados:", db)
	if err := db.Create(&maintenance).Error; err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "failed to create maintenance")
		return
	}

	response.SendSuccess(ctx, http.StatusCreated, maintenance)
}
