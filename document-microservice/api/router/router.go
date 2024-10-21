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

	r.StaticFS("/swagger-docs", http.Dir("/app/document-microservice/api/docs"))
	url := ginSwagger.URL("/swagger-docs/swagger.yml")
	//Document URL: http://localhost:8083/ui-swagger/index.html
	r.GET("/ui-swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	timetable := r.Group("/api/History")
	timetable.Use(middleware.Authentication())
	{
		timetable.GET("/Account/:id", handler.GetHistoryOfVisits()) // for people with his id in jwt
		timetable.GET("/:id", handler.GetHistory())
		timetable.POST("/", handler.CreateNewHistoryVisit())
		timetable.PUT("/:id", handler.UpdateNewHistoryVisit())
	}

	return r
}
