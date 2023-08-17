package projectHandler

import (
	"fmt"
	"time"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateProjectRequest struct {
	Title       string    `json:"title"`
	Client      string    `json:"client"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Status      string    `json:"status"`
	Value       float64   `json:"value"`
	IsPaid      *bool     `json:"isPaid"`
}

type UpdateProjectRequest struct {
	Title       string    `json:"title"`
	Client      string    `json:"client"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Status      string    `json:"status"`
	Value       float64   `json:"value"`
	IsPaid      *bool     `json:"isPaid"`
}

func (req *CreateProjectRequest) ValidateCreateProjectRequest() error {
	if req == nil || req.Title == "" && req.Client == "" && req.Description == "" && req.Deadline.IsZero() && req.Status == "" && req.Value == 0 && req.IsPaid == nil {
		return fmt.Errorf("request body is empty or malformed")
	}
	if req.Title == "" {
		return errParamIsRequired("title", "string")
	}
	if req.Client == "" {
		return errParamIsRequired("client", "string")
	}
	if req.Description == "" {
		return errParamIsRequired("description", "string")
	}
	if req.Deadline.IsZero() {
		return errParamIsRequired("deadline", "time.Time")
	}
	if req.Status == "" {
		return errParamIsRequired("status", "string")
	}
	if req.Value <= 0 {
		return errParamIsRequired("value", "float64")
	}
	return nil
}
