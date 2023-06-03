package myprofile

import (
	"go-hrms-app/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Address struct {
	ID        uint    `gorm:"primarykey"`
	Permanent Contact `json:"permanent"`
	Present   Contact `json:"present"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}

type Contact struct {
	Address string `json:"address"`
	City    string `json:"city"`
	Country string `json:"country"`
}

// // Connect to the MySQL database
// dsn := "user:password@tcp(localhost:3306)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// if err != nil {
// 	log.Fatal(err)
// }

// // Migrate the database schema
// err = db.AutoMigrate(&Address{})
// if err != nil {
// 	log.Fatal(err)
// }

func GetAddress(c *gin.Context) {
	// Query the address record from the database
	var address Address
	result := config.DB.First(&address)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, address)
}

func AddAddress(c *gin.Context) {
	// Parse the request body
	var address Address
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert the address into the database
	result := config.DB.Create(&address)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Address added successfully"})
}
