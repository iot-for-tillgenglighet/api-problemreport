package database

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"

	"github.com/iot-for-tillgenglighet/api-problemreport/pkg/models"
)

var db *gorm.DB

//GetDB returns a pointer to our global database object. Yes, this should be refactored ...
func GetDB() *gorm.DB {
	return db
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

//ConnectToDB extracts connection information from environment variables and
//initiates a connection to the database.
func ConnectToDB() {

	dbHost := os.Getenv("PROBLEMREPORT_DB_HOST")
	username := os.Getenv("PROBLEMREPORT_DB_USER")
	dbName := os.Getenv("PROBLEMREPORT_DB_NAME")
	password := os.Getenv("PROBLEMREPORT_DB_PASSWORD")
	sslMode := getEnv("PROBLEMREPORT_DB_SSLMODE", "require")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", dbHost, username, dbName, sslMode, password)

	for {
		log.Printf("Connecting to database host %s ...\n", dbHost)
		conn, err := gorm.Open("postgres", dbURI)
		if err != nil {
			log.Fatalf("Failed to connect to database %s \n", err)
			time.Sleep(3 * time.Second)
		} else {
			db = conn
			db.Debug().AutoMigrate(&models.ProblemReport{})
			db.Debug().AutoMigrate(&models.ProblemReportCategory{})
			db.Debug().Exec("TRUNCATE TABLE problem_report_categories RESTART IDENTITY; INSERT INTO problem_report_categories(created_at, updated_at, label,report_type,enabled) VALUES(NOW(),NOW(),'Halka','TYPE_ICE', true),(NOW(),NOW(),'Vägskada', 'TYPE_ROAD', true), (NOW(),NOW(),'Otrygghet', 'TYPE_SAFETY', true)")
			return
		}
		defer conn.Close()
	}
}

//Create creates a report
func Create(entity *models.ProblemReport) (*models.ProblemReport, error) {

	currentTime := time.Now().UTC()

	entity.Timestamp = currentTime.Format(time.RFC3339)

	GetDB().Create(entity)

	return entity, nil
}

//GetAll Fetches all problemreports
func GetAll() ([]models.ProblemReport, error) {

	entities := []models.ProblemReport{}
	GetDB().Table("problem_reports").Select("*").Find(&entities)

	return entities, nil
}

//GetCategories fetches all categories
func GetCategories() ([]models.ProblemReportCategory, error) {
	entities := []models.ProblemReportCategory{}
	GetDB().Table("problem_report_categories").Where("enabled = ?", true).Find(&entities)

	return entities, nil
}
