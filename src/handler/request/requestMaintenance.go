package request

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateMaintenanceRequest struct {
	Name string `json:"name"`
	Data string `json:"data"`
	Hora string `json:"hora"`
}

func (r *CreateMaintenanceRequest) Validate() error {
	if r.Name == "" &&
		r.Data == "" &&
		r.Hora == "" {
		return fmt.Errorf("request body is empty")
	}

	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}

	if r.Data == "" {
		return errParamIsRequired("data", "string")
	}

	if r.Hora == "" {
		return errParamIsRequired("hora", "string")
	}

	return nil
}
