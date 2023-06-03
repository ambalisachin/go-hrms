package dashboard

import (
	"go-hrms-app/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Project struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"not null"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`
}

type Circular struct {
	ID    uint      `gorm:"primaryKey"`
	Title string    `gorm:"not null"`
	File  string    `gorm:"not null"`
	Date  time.Time `gorm:"not null"`
}

type Holiday struct {
	ID   uint      `gorm:"primaryKey"`
	Name string    `gorm:"not null"`
	Date time.Time `gorm:"not null"`
}

type Dashboard struct {
	RunningProjects []Project
	Circulars       []Circular
	Holidays        []Holiday
}

// func main() {
// 	dsn := "root:password@tcp(localhost:3306)/hrms?parseTime=True"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// AutoMigrate the Project, Circular, and Holiday structs to create the necessary tables
// 	db.AutoMigrate(&Project{})
// 	db.AutoMigrate(&Circular{})
// 	db.AutoMigrate(&Holiday{})

// 	r := gin.Default()

// r.GET("/dashboard",
func GetDashboard(c *gin.Context) {
	var dashboard Dashboard

	// Fetch running projects
	config.DB.Find(&dashboard.RunningProjects)

	// Fetch circulars
	config.DB.Find(&dashboard.Circulars)

	// Fetch holidays
	config.DB.Find(&dashboard.Holidays)

	c.JSON(http.StatusOK, dashboard)
}
