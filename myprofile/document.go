package myprofile

import (
	"fmt"
	"go-hrms-app/config"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

type Document struct {
	ID        uint   `gorm:"primarykey"`
	FileTitle string `json:"fileTitle"`
	FilePath  string `json:"filePath"`
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
// err = db.AutoMigrate(&Document{})
// if err != nil {
// 	log.Fatal(err)
// }

func GetDocuments(c *gin.Context) {
	// Query the list of documents from the database
	var documents []Document
	result := config.DB.Find(&documents)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, documents)
}

func UploadDocument(c *gin.Context) {
	// Retrieve the uploaded file from the form data
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a unique file name
	fileName := GenerateFileName(file.Filename)

	// Save the uploaded file to the server
	err = c.SaveUploadedFile(file, "uploads/"+fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create a new document record
	newDocument := Document{
		FileTitle: c.PostForm("fileTitle"),
		FilePath:  "uploads/" + fileName,
	}

	// Insert the new document into the database
	result := config.DB.Create(&newDocument)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document uploaded successfully"})
}

func GenerateFileName(originalName string) string {
	ext := filepath.Ext(originalName)
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	return fileName
}
