package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Title       string
	Client      string
	Description string
	Deadline    time.Time
	Status      string
	Value       float64
	IsPaid      bool
}

type ProjectResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Client      string    `json:"client"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Status      string    `json:"status"`
	Value       float64   `json:"value"`
	IsPaid      bool      `json:"isPaid"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"`
}
