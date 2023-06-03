package myprofile

import (
	"go-hrms-app/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Experience struct {
	ID                    uint   `gorm:"primarykey"`
	CompanyName           string `json:"companyName"`
	Position              string `json:"position"`
	WorkDuration          string `json:"workDuration"`
	SalaryCertificate     string `json:"salaryCertificate"`
	ExperienceCertificate string `json:"experienceCertificate"`
	RelievingLetter       string `json:"relievingLetter"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

// // Connect to the MySQL database
// dsn := "user:password@tcp(localhost:3306)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// if err != nil {
// 	log.Fatal(err)
// }

// // Migrate the database schema
// err = db.AutoMigrate(&Experience{})
// if err != nil {
// 	log.Fatal(err)
// }

func GetExperienceList(c *gin.Context) {
	// Query the list of experience records from the database
	var experiences []Experience
	result := config.DB.Find(&experiences)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, experiences)
}

func AddExperience(c *gin.Context) {
	// Parse the request body
	var newExperience Experience
	if err := c.ShouldBindJSON(&newExperience); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert the new experience record into the database
	result := config.DB.Create(&newExperience)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Experience record added successfully"})
}
