package work

import (
	"go-hrms-app/config"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

type Work struct {
	ID          uint   `gorm:"primaryKey"`
	SlNo        uint   `gorm:"not null"`
	Date        string `gorm:"not null"`
	WorkDetails string `gorm:"not null"`
	File        string `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

//err = db.AutoMigrate(&Work{})
// if err != nil {
// 	log.Fatal(err)
// }

func AddWork(c *gin.Context) {
	var work Work
	if err := c.ShouldBindJSON(&work); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the file
	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	work.File = filename
	work.CreatedAt = time.Now()
	work.UpdatedAt = time.Now()

	if err := config.DB.Create(&work).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, work)
}
