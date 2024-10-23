package handler

import (
	"net/http"
	"strconv"
	"strings"
	"wefdzen/internal/database"
	"wefdzen/internal/service"

	"github.com/gin-gonic/gin"
)

func GetInfoAllHospitals() gin.HandlerFunc {
	return func(c *gin.Context) {
		//get data form query
		from := c.Query("from")
		if from == "" {
			from = "0" //если пусто будет по умолчанию 0 => c самого начала
		}
		count := c.Query("count")
		if count == "" {
			count = "10" //если пусто будет по умолчанию 10
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
		hospitals := database.GetListHospitals(userRepo, fromI, countI)
		c.JSON(http.StatusOK, hospitals)
	}
}

func GetInfoHospitalByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idHospital := c.Param("id")
		if _, err := strconv.Atoi(idHospital); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		userRepo := database.NewGormUserRepository()
		hospital := database.GetInfoHospitalByID(userRepo, idHospital)
		if hospital.Name == "" {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
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
		if _, err := strconv.Atoi(idHospital); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		userRepo := database.NewGormUserRepository()
		hospital := database.GetInfoHospitalByID(userRepo, idHospital)

		if hospital.Name == "" {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

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
			c.AbortWithStatus(http.StatusForbidden)
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
			c.AbortWithStatus(http.StatusForbidden)
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
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		//Get ID for soft delete
		idHospital := c.Param("id")
		_, err := strconv.Atoi(idHospital)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

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
		idHospital, err := c.Cookie("idHospital")
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		userRepo := database.NewGormUserRepository()
		hospitalWithRoomExist := database.CheckExistRoomInHospitalID(userRepo, room, idHospital)
		if !hospitalWithRoomExist {
			c.AbortWithStatus(http.StatusNotFound)
		}
		c.AbortWithStatus(http.StatusOK)
	}
}
