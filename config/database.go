package config

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/gorm"
)

var DB *gorm.DB

// Credentials struct can be used to store credentials in a single data type.
type Credentials struct {
	Username string
	Password string
	Server   string
	Dbname   string
}

// Database creates a variable called Database that is set to a Credentials struct.
var Database = Credentials{
	Username: "kadamba",
	Password: "Kadamba@123",
	Server:   "tcp(localhost:3306)",
	Dbname:   "sachindb",
}

// ConnectToDB connects to the database
func (m Credentials) ConnectToDB() *sql.DB {
	dataSourceName := m.Username + ":" + m.Password + "@" + m.Server + "/" + m.Dbname
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour * 1)
	fmt.Println("Connected to DB Successfully....... ")
	return db
}

// NewTable creates new table if the table not exist
func NewTable() {
	db := Database.ConnectToDB()
	defer db.Close()
	//checking for create table for user in db exist or not , if not in db
	//crate a table for user in db
	_, err := db.Query("CREATE TABLE IF NOT EXISTS user(Email varchar(20) NOT NULL, Username varchar(20) NOT NULL, Password varchar(20) NOT NULL)")
	if err != nil {
		fmt.Println(err)

	}
}
