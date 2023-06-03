package leave

import (
	"go-hrms-app/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type LeaveApplication struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	PIN       string
	LeaveType string
	ApplyDate time.Time
	StartDate time.Time
	EndDate   time.Time
	Duration  int
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
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
// 	err = db.AutoMigrate(&LeaveApplication{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Start the HTTP server
// 	err = router.Run(":8080")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func GetLeaveApplication(c *gin.Context) {
	// Get the leave application ID from the request URL
	id := c.Param("id")

	// Query the leave application from the database
	var leave LeaveApplication
	result := config.DB.First(&leave, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, leave)
}

func CreateLeaveApplication(c *gin.Context) {
	// Parse the request body
	var leave LeaveApplication
	if err := c.ShouldBindJSON(&leave); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the created and updated timestamps
	now := time.Now()
	leave.CreatedAt = now
	leave.UpdatedAt = now

	// Insert the leave application into the database
	result := config.DB.Create(&leave)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Leave application created successfully"})
}

func UpdateLeaveApplication(c *gin.Context) {
	// Get the leave application ID from the request URL
	id := c.Param("id")

	// Query the leave application from the database
	var leave LeaveApplication
	result := config.DB.First(&leave, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	// Parse the request body
	var updatedLeave LeaveApplication
	if err := c.ShouldBindJSON(&updatedLeave); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the leave application fields
	leave.LeaveType = updatedLeave.LeaveType
	leave.StartDate = updatedLeave.StartDate
	leave.EndDate = updatedLeave.EndDate
	leave.Duration = updatedLeave.Duration
	leave.Status = updatedLeave.Status
	leave.UpdatedAt = time.Now()

	// Save the updated leave application in the database
	result = config.DB.Save(&leave)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Leave application updated successfully"})
}
