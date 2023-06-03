package leave

import (
	"go-hrms-app/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type LeaveSheet struct {
	ID            uint      `gorm:"primarykey"`
	EmployeeID    string    `json:"employeeID"`
	Name          string    `json:"name"`
	Type          string    `json:"type"`
	LeavesCredit  int       `json:"leavesCredit"`
	LeavesTaken   int       `json:"leavesTaken"`
	LeavesBalance int       `json:"leavesBalance"`
	Date          time.Time `json:"date"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// // Initialize the Gin router
// router := gin.Default()

// // Connect to the MySQL database
// dsn := "user:password@tcp(localhost:3306)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// if err != nil {
// 	log.Fatal(err)
// }

// // Migrate the database schema
// err = db.AutoMigrate(&LeaveSheet{})
// if err != nil {
// 	log.Fatal(err)
// }

// 	// Start the HTTP server
// 	err = router.Run(":8080")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func GetLeaveSheet(c *gin.Context) {
	// Query all leave sheet records from the database
	var leaveSheetList []LeaveSheet
	result := config.DB.Find(&leaveSheetList)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, leaveSheetList)
}

func AddLeaveSheet(c *gin.Context) {
	// Parse the request body
	var leaveSheet LeaveSheet
	if err := c.ShouldBindJSON(&leaveSheet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the created and updated timestamps
	now := time.Now()
	leaveSheet.CreatedAt = now
	leaveSheet.UpdatedAt = now

	// Insert the leave sheet into the database
	result := config.DB.Create(&leaveSheet)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Leave sheet added successfully"})
}

func UpdateLeaveSheet(c *gin.Context) {
	// Get the leave sheet ID from the request URL
	id := c.Param("id")

	// Query the leave sheet from the database
	var leaveSheet LeaveSheet
	result := config.DB.First(&leaveSheet, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	// Parse the request body
	var updatedLeaveSheet LeaveSheet
	if err := c.ShouldBindJSON(&updatedLeaveSheet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the leave sheet fields
	leaveSheet.LeavesCredit = updatedLeaveSheet.LeavesCredit
	leaveSheet.LeavesTaken = updatedLeaveSheet.LeavesTaken
	leaveSheet.LeavesBalance = updatedLeaveSheet.LeavesBalance
	leaveSheet.UpdatedAt = time.Now()

	// Save the updated leave sheet in the database
	result = config.DB.Save(&leaveSheet)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Leave sheet updated successfully"})
}
