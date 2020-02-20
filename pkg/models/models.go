package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type ProblemReport struct {
	gorm.Model
	Latitude  float64
	Longitude float64
	Type      string
	Timestamp string `gorm:"unique_index:idx_device_timestamp"`
}
