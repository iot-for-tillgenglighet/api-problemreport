package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//ProblemReport Base type for problem report
type ProblemReport struct {
	gorm.Model
	Latitude  float64
	Longitude float64
	Type      string
	Timestamp string `gorm:"unique_index:idx_device_timestamp"`
}

//ProblemReportCategory Base object for problem report category
type ProblemReportCategory struct {
	gorm.Model
	Label      string
	ReportType string
	Enabled    bool
}
