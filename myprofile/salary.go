package myprofile

import (
	"go-hrms-app/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Salary struct {
	ID          uint    `gorm:"primarykey"`
	SalaryType  string  `json:"salaryType"`
	TaxType     string  `json:"taxType"`
	Salary      float64 `json:"salary"`
	Tax         float64 `json:"tax"`
	TDS         float64 `json:"tds"`
	TotalSalary float64 `json:"totalSalary"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Connect to the MySQL database
// dsn := "user:password@tcp(localhost:3306)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// if err != nil {
// 	log.Fatal(err)
// }

// // Migrate the database schema
// err = db.AutoMigrate(&Salary{})
// if err != nil {
// 	log.Fatal(err)
// }

func GetSalaries(c *gin.Context) {
	// Query the list of salaries from the database
	var salaries []Salary
	result := config.DB.Find(&salaries)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, salaries)
}

func AddSalary(c *gin.Context) {
	// Parse the request body
	var newSalary Salary
	if err := c.ShouldBindJSON(&newSalary); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Calculate the total salary based on salary and tax
	newSalary.TotalSalary = newSalary.Salary - newSalary.Tax - newSalary.TDS

	// Insert the new salary into the database
	result := config.DB.Create(&newSalary)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Salary added successfully"})
}
