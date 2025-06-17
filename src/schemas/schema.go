package schemas

import "gorm.io/gorm"

type Maintenance struct {
	gorm.Model
	Name string
	Data string
	Hora string
}

type MaintenanceResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Data string `json:"data"`
	Hora string `json:"hora"`
}