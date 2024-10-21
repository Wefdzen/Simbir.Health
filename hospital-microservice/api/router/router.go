package router

import (
	"net/http"
	"wefdzen/internal/handler"
	"wefdzen/internal/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.StaticFS("/swagger-docs", http.Dir("/app/hospital-microservice/api/docs"))
	url := ginSwagger.URL("/swagger-docs/swagger.yml")
	//Hospital URL: http://localhost:8081/ui-swagger/index.html
	r.GET("/ui-swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

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
