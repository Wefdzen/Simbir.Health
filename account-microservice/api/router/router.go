package router

import (
	"wefdzen/internal/handler"
	"wefdzen/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	userAuth := r.Group("/api/Authentication")
	{
		userAuth.POST("/SignUp", handler.SignUp())
		userAuth.POST("/SignIn", handler.SignIn())
		userAuth.PUT("/SignOut", middleware.Authentication(), handler.SignOut()) //for auth users
		userAuth.GET("/Validate", handler.ValidateToken())                       //интроспекция токена
		userAuth.POST("/Refresh", handler.Refresh())

	}
	//TODO role amind authoriz middleware
	userAccounts := r.Group("/api/Accounts")
	userAccounts.Use(middleware.Authentication())
	{
		userAccounts.GET("/Me", handler.Me())
		userAccounts.PUT("/Update", handler.UpdateDataByUser())
		userAccounts.GET("/", handler.GetAllAccounts())       //Role: admin for get all acc
		userAccounts.POST("/", handler.CreateAccount())       //Role: admin for create acc with role
		userAccounts.PUT("/:id", handler.UpdateDataByAdmin()) //Role: admin for change data of acc
		userAccounts.DELETE("/:id", handler.DeleteAccount())  //Role: admin for delete soft account
	}

	userDoctor := r.Group("/api/Doctors") // Role: user and >
	userDoctor.Use(middleware.Authentication())
	{
		userDoctor.GET("/", handler.GetAllDoctors())     // get list of doctors
		userDoctor.GET("/:id", handler.GetInforDoctor()) // get info about doctor
	}

	return r
}
