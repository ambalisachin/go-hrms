package myprofile

import (
	"go-hrms-app/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Deputation struct {
	gorm.Model
	PromotedDate        time.Time
	PreviousDepartment  string
	PreviousDesignation string
	CurrentDepartment   string
	CurrentDesignation  string
}

// Auto migrate the Deputation model
//db.AutoMigrate(&Deputation{})

func AddDeputation(c *gin.Context) {
	var deputation Deputation

	if err := c.ShouldBindJSON(&deputation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	result := config.DB.Create(&deputation)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create deputation"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Deputation created successfully"})
}

func GetDeputation(c *gin.Context) {
	id := c.Param("id")

	var deputation Deputation
	result := config.DB.First(&deputation, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Deputation not found"})
		return
	}

	c.JSON(http.StatusOK, deputation)
}
