package myprofile

import (
	"go-hrms-app/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Leave struct {
	ID        uint   `gorm:"primarykey"`
	LeaveType string `json:"leaveType"`
	Days      int    `json:"days"`
	Year      int    `json:"year"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Connect to the MySQL database
// dsn := "user:password@tcp(localhost:3306)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// if err != nil {
// 	log.Fatal(err)
// }

// // Migrate the database schema
// err = db.AutoMigrate(&Leave{})
// if err != nil {
// 	log.Fatal(err)
// }

// Routes
// router.GET("/leaves", getLeaves)
// router.POST("/leaves", addLeave)

func GetLeaves(c *gin.Context) {
	// Query the list of leaves from the database
	var leaves []Leave
	result := config.DB.Find(&leaves)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, leaves)
}

func AddLeave(c *gin.Context) {
	// Parse the request body
	var newLeave Leave
	if err := c.ShouldBindJSON(&newLeave); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert the new leave into the database
	result := config.DB.Create(&newLeave)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Leave added successfully"})
}
