package handler

import (
	"net/http"
	"strconv"
	"wefdzen/internal/database"

	"github.com/gin-gonic/gin"
)

type DoctorResponse struct {
	FullName string
}

func GetAllDoctors() gin.HandlerFunc {
	return func(c *gin.Context) {
		nameFilter := c.Query("nameFilter")
		from := c.Query("from")
		count := c.Query("count")
		fromI, _ := strconv.Atoi(from)
		countI, _ := strconv.Atoi(count)

		userRepo := database.NewGormUserRepository()
		doctorsWithPar := database.GetFullNameHowDoctors(userRepo, fromI, countI, nameFilter)
		var responseSlice []DoctorResponse
		for _, user := range doctorsWithPar {
			responseSlice = append(responseSlice, DoctorResponse{
				FullName: user.LastName + " " + user.FirstName,
			})
		}
		c.JSON(http.StatusOK, responseSlice)
	}
}

func GetInforDoctor() gin.HandlerFunc {
	return func(c *gin.Context) {
		idUser := c.Param("id")
		userRepo := database.NewGormUserRepository()
		doctorAllInfo := database.GetInfoIDDoctor(userRepo, idUser)
		var doctor DoctorResponse
		doctor.FullName = doctorAllInfo.LastName + " " + doctorAllInfo.FirstName
		c.JSON(http.StatusOK, doctor)
	}
}
