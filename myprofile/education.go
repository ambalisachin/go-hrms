package myprofile

import (
	"go-hrms-app/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Education struct {
	ID          uint   `gorm:"primarykey"`
	Certificate string `json:"certificate"`
	Institute   string `json:"institute"`
	File        string `json:"file"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Connect to the MySQL database
// dsn := "user:password@tcp(localhost:3306)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// if err != nil {
// 	log.Fatal(err)
// }

// // Migrate the database schema
// err = db.AutoMigrate(&Education{})
// if err != nil {
// 	log.Fatal(err)
// }

func GetEducationList(c *gin.Context) {
	// Query the list of education records from the database
	var educations []Education
	result := config.DB.Find(&educations)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, educations)
}

func AddEducation(c *gin.Context) {
	// Parse the request body
	var newEducation Education
	if err := c.ShouldBindJSON(&newEducation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert the new education record into the database
	result := config.DB.Create(&newEducation)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Education record added successfully"})
}
