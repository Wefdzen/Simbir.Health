package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"wefdzen/internal/database"
	"wefdzen/internal/service"

	"github.com/gin-gonic/gin"
)

type HospitalsResponse struct {
	Name string
}

func GetInfoAllHospitals() gin.HandlerFunc {
	return func(c *gin.Context) {
		//get data form query
		from := c.Query("from")
		count := c.Query("count")
		fromI, _ := strconv.Atoi(from)
		countI, _ := strconv.Atoi(count)

		userRepo := database.NewGormUserRepository()
		hospitals := database.GetListHospitals(userRepo, fromI, countI)
		var responseSlice []HospitalsResponse
		for _, hospital := range hospitals {
			responseSlice = append(responseSlice, HospitalsResponse{
				Name: hospital.Name,
			})
		}
		c.JSON(http.StatusOK, responseSlice)
	}
}

func GetInfoHospitalByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idHospital := c.Param("id")
		userRepo := database.NewGormUserRepository()
		hospital := database.GetInfoHospitalByID(userRepo, idHospital)
		c.JSON(http.StatusOK, gin.H{
			"name":         hospital.Name,
			"address":      hospital.Address,
			"contactPhone": hospital.ContactPhone,
		})
	}
}

func GetListRoomsInHospitalByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idHospital := c.Param("id")
		userRepo := database.NewGormUserRepository()
		hospital := database.GetInfoHospitalByID(userRepo, idHospital)
		c.JSON(http.StatusOK, gin.H{
			"rooms": hospital.Rooms,
		})

	}
}

func CreateNewHospitalByAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		//only for admin
		roles := service.Authorization(c)
		if !strings.Contains(roles, "admin") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		var jsonInput database.Hospital
		if err := c.BindJSON(&jsonInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		//connect to db
		userRepo := database.NewGormUserRepository()
		database.CreateHospitalByAdmin(userRepo, &jsonInput)
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})

	}
}

func UpdateHospitalByAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		//only for admin
		roles := service.Authorization(c)
		if !strings.Contains(roles, "admin") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		//get data from request
		var jsonInput database.Hospital
		if err := c.BindJSON(&jsonInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		//get id for udpate date
		idHospital := c.Param("id")
		userRepo := database.NewGormUserRepository()
		database.UpdateHospitalData(userRepo, idHospital, &jsonInput)
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}

func SoftDeleteHospitalByAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		//only for admin
		roles := service.Authorization(c)
		if !strings.Contains(roles, "admin") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		//Get ID for soft delete
		idHospital := c.Param("id")
		userRepo := database.NewGormUserRepository()
		database.SoftDeleteByAdmin(userRepo, idHospital)
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}

func CheckExistRoomInHospital() gin.HandlerFunc {
	return func(c *gin.Context) {
		room, err := c.Cookie("room")
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		fmt.Println("ROOOOOOOOOOOOOm: ", room)
		idHospital, err := c.Cookie("idHospital")
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		userRepo := database.NewGormUserRepository()
		hospitalWithRoomExist := database.CheckExistRoomInHospitalID(userRepo, room, idHospital)
		if !hospitalWithRoomExist {
			c.AbortWithStatus(http.StatusBadRequest)
		}
		c.AbortWithStatus(http.StatusOK)
	}
}
