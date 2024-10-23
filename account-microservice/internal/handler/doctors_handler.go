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
		if from == "" {
			from = "0" //by default 0
		}
		count := c.Query("count")
		if count == "" {
			count = "10" //by default 10
		}
		fromI, err := strconv.Atoi(from)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		countI, err := strconv.Atoi(count)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		userRepo := database.NewGormUserRepository()
		doctorsWithPar := database.GetFullNameHowDoctors(userRepo, fromI, countI, nameFilter)
		var responseSlice []DoctorResponse
		for _, user := range doctorsWithPar {
			responseSlice = append(responseSlice, DoctorResponse{
				FullName: user.LastName + " " + user.FirstName,
			})
		}
		//значит записи небыли найдены
		if responseSlice[0].FullName == " " { // FulltName contain "" + " " + ""
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusOK, responseSlice)
	}
}

func GetInforDoctor() gin.HandlerFunc {
	return func(c *gin.Context) {
		idDoctor := c.Param("id")
		userRepo := database.NewGormUserRepository()
		doctorAllInfo := database.GetInfoIDDoctor(userRepo, idDoctor)
		var doctor DoctorResponse
		doctor.FullName = doctorAllInfo.LastName + " " + doctorAllInfo.FirstName

		//значит записи небылo
		doctorExist := database.CheckExistDoctorID(userRepo, idDoctor)
		if !doctorExist {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusOK, doctor)
	}
}

func CheckExistDoctor() gin.HandlerFunc {
	return func(c *gin.Context) {
		idDoctor, err := c.Cookie("idDoctor")
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		userRepo := database.NewGormUserRepository()
		doctorExist := database.CheckExistDoctorID(userRepo, idDoctor)
		if !doctorExist {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.AbortWithStatus(http.StatusOK)
	}
}
