package models

import (
	"gorm.io/gorm"
)

//ProblemReport Base type for problem report
type ProblemReport struct {
	gorm.Model
	Latitude  float64
	Longitude float64
	Type      string
	Timestamp string
}

//ProblemReportCategory Base object for problem report category
type ProblemReportCategory struct {
	gorm.Model
	Label      string
	ReportType string
	Enabled    bool
}
