package attendance

import (
	"go-hrms-app/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AttendanceList struct {
	ID          uint      `gorm:"primarykey"`
	Employee    string    `json:"employee"`
	Date        time.Time `json:"date"`
	SignInTime  time.Time `json:"signInTime"`
	SignOutTime time.Time `json:"signOutTime"`
	WorkedHours float64   `json:"workedHours"`
	Action      string    `json:"action"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// func main() {
// 	// Initialize the Gin router
// 	router := gin.Default()

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

// 	// Start the HTTP server
// 	err = router.Run(":8080")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func GetAttendanceList(c *gin.Context) {
	// Query all attendance records from the database
	var attendanceList []AttendanceList
	result := config.DB.Find(&attendanceList)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, attendanceList)
}

// func AddAttendance(c *gin.Context) {
// 	// Parse the request body
// 	var attendance Attendance
// 	if err := c.ShouldBindJSON(&attendance); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Set the created and updated timestamps
// 	now := time.Now()
// 	attendance.CreatedAt = now
// 	attendance.UpdatedAt = now

// 	// Insert the attendance into the database
// 	result := db.Create(&attendance)
// 	if result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Attendance added successfully"})
// }

func SignOutAttendance(c *gin.Context) {
	// Get the attendance ID from the request URL
	id := c.Param("id")

	// Query the attendance from the database
	var attendance AttendanceList
	result := config.DB.First(&attendance, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	// Parse the request body
	var SignOutAttendance AttendanceList
	if err := c.ShouldBindJSON(&SignOutAttendance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the attendance fields
	attendance.SignOutTime = SignOutAttendance.SignOutTime
	attendance.WorkedHours = SignOutAttendance.WorkedHours
	attendance.Action = SignOutAttendance.Action
	attendance.UpdatedAt = time.Now()

	// Save the updated attendance in the database
	result = config.DB.Save(&attendance)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Attendance signout successfully"})
}
