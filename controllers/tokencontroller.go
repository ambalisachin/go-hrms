package controllers

import (
	"fmt"
	"go-hrms-app/config"
	models "go-hrms-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	config.NewTable()
	db := config.Database.ConnectToDB()
	defer db.Close()
	_, err := db.Query("insert into user(Email varchar(20) NOT NULL, Username varchar(20) NOT NULL, Password varchar(20) NOT NULL)")
	if err != nil {
		fmt.Println(err)

	}

	record := config.DB.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}
