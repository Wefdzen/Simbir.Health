package router

import (
	"wefdzen/internal/handler"
	"wefdzen/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	hospitals := r.Group("/api/Hospitals")
	hospitals.Use(middleware.Authentication())
	{
		hospitals.GET("/", handler.GetInfoAllHospitals()) // get list of hospitals
		hospitals.GET("/:id", handler.GetInfoHospitalByID())
		hospitals.GET("/:id/Rooms", handler.GetListRoomsInHospitalByID())
		hospitals.POST("/", handler.CreateNewHospitalByAdmin())       // for admin create hospital
		hospitals.PUT("/:id", handler.UpdateHospitalByAdmin())        //for admin update info about hospit by id
		hospitals.DELETE("/:id", handler.SoftDeleteHospitalByAdmin()) //for admin SOFT delete hospit by id
		hospitals.GET("/Exist", handler.CheckExistRoomInHospital())
	}

	return r
}
