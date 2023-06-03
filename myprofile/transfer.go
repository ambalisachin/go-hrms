package myprofile

import (
	"go-hrms-app/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Transfer struct {
	gorm.Model
	TransferDate   time.Time
	PreviousBranch string
	CurrentBranch  string
}

// Auto migrate the Transfer model
//db.AutoMigrate(&Transfer{})

func AddTransfer(c *gin.Context) {
	var transfer Transfer

	if err := c.ShouldBindJSON(&transfer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	result := config.DB.Create(&transfer)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transfer"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Transfer created successfully"})
}

func GetTransfer(c *gin.Context) {
	id := c.Param("id")

	var transfer Transfer
	result := config.DB.First(&transfer, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transfer not found"})
		return
	}

	c.JSON(http.StatusOK, transfer)
}
