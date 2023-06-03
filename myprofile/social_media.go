package myprofile

import (
	"go-hrms-app/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name         string
	FacebookURL  string
	TwitterURL   string
	LinkedInURL  string
	InstagramURL string
}

// Auto migrate the Employee model
//db.AutoMigrate(&Employee{})

func GetSocialMedia(c *gin.Context) {
	id := c.Param("id")

	var employee Employee
	result := config.DB.First(&employee, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":          employee.Name,
		"facebook_url":  employee.FacebookURL,
		"twitter_url":   employee.TwitterURL,
		"linkedin_url":  employee.LinkedInURL,
		"instagram_url": employee.InstagramURL,
	})
}

func CreatSocialMedia(c *gin.Context) {
	id := c.Param("id")

	var employee Employee
	result := config.DB.First(&employee, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	var payload struct {
		FacebookURL  string `json:"facebook_url"`
		TwitterURL   string `json:"twitter_url"`
		LinkedInURL  string `json:"linkedin_url"`
		InstagramURL string `json:"instagram_url"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	employee.FacebookURL = payload.FacebookURL
	employee.TwitterURL = payload.TwitterURL
	employee.LinkedInURL = payload.LinkedInURL
	employee.InstagramURL = payload.InstagramURL

	config.DB.Save(&employee)

	c.JSON(http.StatusOK, gin.H{"message": "Social media URLs updated successfully"})
}
