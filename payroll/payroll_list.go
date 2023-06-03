package payroll

import (
	"go-hrms-app/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Payroll struct {
	ID        uint      `gorm:"primaryKey"`
	Month     string    `gorm:"not null"`
	Salary    float64   `gorm:"not null"`
	Loan      float64   `gorm:"not null"`
	TotalPaid float64   `gorm:"not null"`
	PayDate   time.Time `gorm:"not null"`
	Status    string    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// err = db.AutoMigrate(&Payroll{})
// if err != nil {
// 	log.Fatal(err)
// }

func GetPayrollList(c *gin.Context) {
	var payroll []Payroll
	if err := config.DB.Find(&payroll).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payroll)
}

func AddPayrollList(c *gin.Context) {
	var payroll Payroll
	if err := c.ShouldBindJSON(&payroll); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payroll.TotalPaid = payroll.Salary - payroll.Loan
	payroll.PayDate = time.Now()
	payroll.CreatedAt = time.Now()
	payroll.UpdatedAt = time.Now()

	if err := config.DB.Create(&payroll).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payroll)
}
