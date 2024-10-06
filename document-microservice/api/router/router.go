package router

import (
	"wefdzen/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	timetable := r.Group("/api/History")
	timetable.Use(middleware.Authentication())
	{
		timetable.GET("/Account/:id") // for people with his id in jwt
		timetable.GET("/:id")
		timetable.POST("/")
		timetable.PUT(":id")
	}

	return r
}
