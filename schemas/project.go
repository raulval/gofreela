package schemas

import (
	"gorm.io/gorm"
)

type Project struct {
		gorm.Model
    Title           string
    Client          string
    Description     string
    Deadline        string
    Status          string
    Value           float64
    IsPaid 					bool
}