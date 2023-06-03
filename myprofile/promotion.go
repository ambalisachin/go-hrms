package myprofile

import (
	"go-hrms-app/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Promotion struct {
	gorm.Model
	PromotedDate        time.Time
	PreviousDepartment  string
	PreviousDesignation string
	CurrentDepartment   string
	CurrentDesignation  string
}

// Auto migrate the Promotion model
//	db.AutoMigrate(&Promotion{})

func AddPromotion(c *gin.Context) {
	var promotion Promotion

	if err := c.ShouldBindJSON(&promotion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	result := config.DB.Create(&promotion)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create promotion"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Promotion created successfully"})
}

func GetPromotion(c *gin.Context) {
	id := c.Param("id")

	var promotion Promotion
	result := config.DB.First(&promotion, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Promotion not found"})
		return
	}

	c.JSON(http.StatusOK, promotion)
}
