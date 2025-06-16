package handler

type CreateMaintenanceRequest struct {
	Name string `json: "name"`
	Data string `json: "data"`
	Hora string `json: "hora"`
}

func (r *CreateMaintenanceRequest) Validate() error {
	
}
