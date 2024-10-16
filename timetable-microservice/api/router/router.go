package router

import (
	"wefdzen/internal/handler"
	"wefdzen/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	timetable := r.Group("/api/Timetable")
	timetable.Use(middleware.Authentication())
	{
		timetable.POST("/", handler.CreateRecordInTimetable())
		timetable.PUT("/:id", handler.UpdateRecordInTimetable())
		timetable.DELETE("/:id", handler.DeleteRecordFromTimetable())
		timetable.DELETE("/Doctor/:id", handler.DeleteTimetableForDoctor())
		timetable.DELETE("/Hospital/:id", handler.DeleteTimetableForHospital())
		timetable.GET("/Hospital/:id", handler.GetTimetableByIdHospital())
		timetable.GET("/Doctor/:id", handler.GetTimetableByIdDoctor())
		timetable.GET("/Hospital/:id/Room/:room", handler.GetTimetableInHospitalRoom())
		timetable.GET("/:id/Appointments", handler.GetFreeAppointments())
		timetable.POST("/:id/Appointments", handler.RecordingToAppointment())
	}

	timetable2 := r.Group("/api/Appointment")
	timetable2.Use(middleware.Authentication())
	{
		timetable2.DELETE("/:id", handler.CancelingAppointment())
	}
	return r
}
