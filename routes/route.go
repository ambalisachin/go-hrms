package routes

import (
	"go-hrms-app/attendance"
	"go-hrms-app/controllers"
	"go-hrms-app/dashboard"
	"go-hrms-app/leave"
	"go-hrms-app/middlewares"
	"go-hrms-app/myprofile"
	"go-hrms-app/payroll"
	"go-hrms-app/work"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{

		//api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		//api.POST("/logout", controllers.AuthMiddleware(), controllers.Logout)

		api.POST("/login", controllers.Login)
		api.GET("/leavesheet", leave.GetLeaveSheet)
		api.POST("/leavesheet", leave.AddLeaveSheet)
		api.PUT("/leavesheet/:id", leave.UpdateLeaveSheet)

		api.GET("/leave/:id", leave.GetLeaveApplication)
		api.POST("/leave", leave.CreateLeaveApplication)
		api.PUT("/leave/:id", leave.UpdateLeaveApplication)
		api.POST("/attendance", attendance.AddAttendance)
		api.GET("/attendance", attendance.GetAttendanceList)
		api.POST("/attendace/:id", attendance.SignOutAttendance)
		api.GET("/dashboard", dashboard.GetDashboard)
		api.GET("/personalinfo", myprofile.GetPersonalInfo)
		api.POST("/personalinfo", myprofile.AddPersonalInfo)
		api.PUT("/personalinfo/:id", myprofile.UpdatePersonalInfo)
		api.GET("/address", myprofile.GetAddress)
		api.POST("/address", myprofile.AddAddress)
		api.GET("/education", myprofile.GetEducationList)
		api.POST("/education", myprofile.AddEducation)
		api.GET("/experience", myprofile.GetExperienceList)
		api.POST("/experience", myprofile.AddExperience)
		api.GET("/bank-account", myprofile.GetBankAccounts)
		api.POST("/bank-account", myprofile.AddBankAccount)
		api.GET("/bank-account/:ifscCode", myprofile.GetBankAccountByIFSC)
		api.GET("/documents", myprofile.GetDocuments)
		api.POST("/documents", myprofile.UploadDocument)
		api.GET("/salaries", myprofile.GetSalaries)
		api.POST("/salaries", myprofile.AddSalary)
		api.GET("/leaves", myprofile.GetLeaves)
		api.POST("/leaves", myprofile.AddLeave)
		api.GET("/employee/:id", myprofile.GetSocialMedia)
		api.POST("/employeei/:d", myprofile.CreatSocialMedia)
		api.GET("/promotions/:id", myprofile.GetPromotion)
		api.POST("/promotions", myprofile.AddPromotion)
		api.GET("/transfers/:id", myprofile.GetTransfer)
		api.POST("/transfers", myprofile.AddTransfer)
		api.GET("/deputations/:id", myprofile.GetDeputation)
		api.POST("/deputations", myprofile.AddDeputation)
		api.GET("/payroll/", payroll.GetPayrollList)
		api.POST("/payroll", payroll.AddPayrollList)
		api.POST("/work", work.AddWork)

		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
