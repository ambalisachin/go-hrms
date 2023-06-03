package myprofile

import (
	"go-hrms-app/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type BankAccount struct {
	ID             uint   `gorm:"primarykey"`
	IfscCode       string `json:"ifscCode"`
	BankName       string `json:"bankName"`
	BranchName     string `json:"branchName"`
	BankHolderName string `json:"bankHolderName"`
	AccountNumber  string `json:"accountNumber"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// Connect to the MySQL database
///dsn := "user:password@tcp(localhost:3306)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// if err != nil {
// 	log.Fatal(err)
// }

// Migrate the database schema
// _,err = db.AutoMigrate(&BankAccount{})
// if err != nil {
// 	log.Fatal(err)
// }

func GetBankAccounts(c *gin.Context) {
	// Query the list of bank accounts from the database
	var bankAccounts []BankAccount
	result := config.DB.Find(&bankAccounts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, bankAccounts)
}

func AddBankAccount(c *gin.Context) {
	// Parse the request body
	var newBankAccount BankAccount
	if err := c.ShouldBindJSON(&newBankAccount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert the new bank account into the database
	result := config.DB.Create(&newBankAccount)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bank account added successfully"})
}

func GetBankAccountByIFSC(c *gin.Context) {
	// Get the IFSC code from the request parameters
	ifscCode := c.Param("ifscCode")

	// Query the bank account from the database by IFSC code
	var bankAccount BankAccount
	result := config.DB.Where("ifsc_code = ?", ifscCode).First(&bankAccount)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bank account not found"})
		return
	}

	c.JSON(http.StatusOK, bankAccount)
}
