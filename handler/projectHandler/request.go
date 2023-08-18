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
	Title       string  `json:"title"`
	Client      string  `json:"client"`
	Description string  `json:"description"`
	Deadline    string  `json:"deadline"`
	Status      string  `json:"status"`
	Value       float64 `json:"value"`
	IsPaid      *bool   `json:"isPaid"`
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
		return fmt.Errorf("deadline is required and must be in the format ISO 8601")
	}
	if req.Deadline.Before(time.Now()) {
		return fmt.Errorf("deadline cannot be in the past")
	}
	if req.Status == "" {
		return errParamIsRequired("status", "string")
	}
	if req.Value <= 0 {
		return errParamIsRequired("value", "float64")
	}
	return nil
}

func (req *UpdateProjectRequest) ValidateUpdateProjectRequest() error {
	fmt.Println(req.Deadline)
	if req == nil || req.Title == "" && req.Client == "" && req.Description == "" && req.Deadline == "" && req.Status == "" && req.Value == 0 && req.IsPaid == nil {
		return fmt.Errorf("at least one valid field must be provided")
	}
	return nil
}

func ParseDeadline(deadline string) (time.Time, error) {
	if deadline == "" {
		return time.Time{}, nil
	}

	deadlineTime, err := time.Parse(time.RFC3339, deadline)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid format for deadline, it must be in the format ISO 8601")
	}

	if deadlineTime.Before(time.Now()) {
		return time.Time{}, fmt.Errorf("deadline cannot be in the past")
	}

	return deadlineTime, nil
}
