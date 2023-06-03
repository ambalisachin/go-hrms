package attendance

import (
	"go-hrms-app/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Attendance struct {
	ID         uint      `gorm:"primarykey"`
	Employee   string    `json:"employee"`
	Date       time.Time `json:"date"`
	SignInTime time.Time `json:"signInTime"`
	Place      string    `json:"place"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// 	// Connect to the MySQL database
// 	dsn := "user:password@tcp(localhost:3306)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Migrate the database schema
// 	err = db.AutoMigrate(&Attendance{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

func AddAttendance(c *gin.Context) {
	// Parse the request body
	var attendance Attendance
	if err := c.ShouldBindJSON(&attendance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the created and updated timestamps
	now := time.Now()
	attendance.CreatedAt = now
	attendance.UpdatedAt = now

	// Insert the attendance into the database
	result := config.DB.Create(&attendance)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Attendance added successfully"})
}
