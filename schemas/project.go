package schemas

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Project struct {
	ID          uuid.UUID      `gorm:"primaryKey,type:uuid" json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	Title       string         `gorm:"not null" json:"title"`
	Client      string         `gorm:"not null" json:"client"`
	Description string         `json:"description"`
	Deadline    time.Time      `json:"deadline"`
	Status      string         `gorm:"not null" json:"status"`
	Value       float64        `gorm:"not null" json:"value"`
	IsPaid      bool           `gorm:"not null" json:"isPaid"`
}

type ProjectResponse struct {
	ID          uuid.UUID `json:"id"`
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
