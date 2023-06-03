package myprofile

import (
	"go-hrms-app/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PersonalInfo struct {
	ID              uint      `gorm:"primarykey"`
	Pin             string    `json:"pin"`
	FirstName       string    `json:"firstName"`
	LastName        string    `json:"lastName"`
	BloodGroup      string    `json:"bloodGroup"`
	Gender          string    `json:"gender"`
	DateOfBirth     time.Time `json:"dateOfBirth"`
	ContactNumber   string    `json:"contactNumber"`
	Branch          string    `json:"branch"`
	DateOfJoining   time.Time `json:"dateOfJoining"`
	ContractEndDate time.Time `json:"contractEndDate"`
	Email           string    `json:"email"`
	AdharNumber     string    `json:"adharNumber"`
	PANNumber       string    `json:"panNumber"`
	EmergencyName   string    `json:"emergencyName"`
	EmergencyPhone  string    `json:"emergencyPhone"`
	AdharUpload     string    `json:"adharUpload"`
	PANUpload       string    `json:"panUpload"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Connect to the MySQL database
// dsn := "user:password@tcp(localhost:3306)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// if err != nil {
// 	log.Fatal(err)
// }

// // Migrate the database schema
// err = db.AutoMigrate(&PersonalInfo{})
// if err != nil {
// 	log.Fatal(err)
// }

// 	// Start the HTTP server
// 	err = router.Run(":8080")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func GetPersonalInfo(c *gin.Context) {
	// Query all personal info records from the database
	var personalInfoList []PersonalInfo
	result := config.DB.Find(&personalInfoList)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, personalInfoList)
}

func AddPersonalInfo(c *gin.Context) {
	// Parse the request body
	var personalInfo PersonalInfo
	if err := c.ShouldBindJSON(&personalInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the created and updated timestamps
	now := time.Now()
	personalInfo.CreatedAt = now
	personalInfo.UpdatedAt = now

	// Insert the personal info into the database
	result := config.DB.Create(&personalInfo)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Personal info added successfully"})
}

func UpdatePersonalInfo(c *gin.Context) {
	// Get the personal info ID from the request URL
	id := c.Param("id")

	// Query the personal info from the database
	var personalInfo PersonalInfo
	result := config.DB.First(&personalInfo, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	// Parse the request body
	var updatedPersonalInfo PersonalInfo
	if err := c.ShouldBindJSON(&updatedPersonalInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the personal info fields
	personalInfo.Pin = updatedPersonalInfo.Pin
	personalInfo.FirstName = updatedPersonalInfo.FirstName
	personalInfo.LastName = updatedPersonalInfo.LastName
	personalInfo.BloodGroup = updatedPersonalInfo.BloodGroup
	personalInfo.Gender = updatedPersonalInfo.Gender
	personalInfo.DateOfBirth = updatedPersonalInfo.DateOfBirth
	personalInfo.ContactNumber = updatedPersonalInfo.ContactNumber
	personalInfo.Branch = updatedPersonalInfo.Branch
	personalInfo.DateOfJoining = updatedPersonalInfo.DateOfJoining
	personalInfo.ContractEndDate = updatedPersonalInfo.ContractEndDate
	personalInfo.Email = updatedPersonalInfo.Email
	personalInfo.AdharNumber = updatedPersonalInfo.AdharNumber
	personalInfo.PANNumber = updatedPersonalInfo.PANNumber
	personalInfo.EmergencyName = updatedPersonalInfo.EmergencyName
	personalInfo.EmergencyPhone = updatedPersonalInfo.EmergencyPhone
	personalInfo.AdharUpload = updatedPersonalInfo.AdharUpload
	personalInfo.PANUpload = updatedPersonalInfo.PANUpload
	personalInfo.UpdatedAt = time.Now()

	// Save the updated personal info in the database
	result = config.DB.Save(&personalInfo)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Personal info updated successfully"})
}
