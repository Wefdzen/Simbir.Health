package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"wefdzen/internal/database"
	"wefdzen/internal/service"

	"github.com/gin-gonic/gin"
)

func CreateRecordInTimetable() gin.HandlerFunc {
	return func(c *gin.Context) {
		//only for admin or manager
		roles := service.Authorization(c)
		if !strings.Contains(roles, "admin") && !strings.Contains(roles, "manager") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		//get data from request
		var jsonInput database.Timetable
		if err := c.BindJSON(&jsonInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		existDoctor, err := service.CheckExistDoctor(strconv.Itoa(jsonInput.DoctorId), c)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		existRoomInHospital, err := service.CheckExistRoomInHospital(jsonInput.Room, strconv.Itoa(jsonInput.HospitalId), c)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		//Checking paramaters time in body from req
		ans := service.CheckMultiplyOfTime(jsonInput.From, jsonInput.To, c)
		if !ans {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		ans = service.CheckTimeDifference(jsonInput.From, jsonInput.To)
		if !ans {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if existDoctor && existRoomInHospital { // если все прошло хорошо будет создано новое расписание
			userRepo := database.NewGormUserRepository()
			database.CreateTimetableForDoctor(userRepo, &jsonInput)
			c.AbortWithStatus(http.StatusOK) //mb change
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
func UpdateRecordInTimetable() gin.HandlerFunc {
	return func(c *gin.Context) {
		//only for admin or manager
		roles := service.Authorization(c)
		if !strings.Contains(roles, "admin") && !strings.Contains(roles, "manager") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		idTimetable := c.Param("id")
		_, err := strconv.Atoi(idTimetable)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		//get data from request
		var jsonInput database.Timetable
		if err := c.BindJSON(&jsonInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		existDoctor, err := service.CheckExistDoctor(strconv.Itoa(jsonInput.DoctorId), c)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		existRoomInHospital, err := service.CheckExistRoomInHospital(jsonInput.Room, strconv.Itoa(jsonInput.HospitalId), c)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		//Checking paramaters time in body from req
		service.CheckMultiplyOfTime(jsonInput.From, jsonInput.To, c)
		if ans := service.CheckTimeDifference(jsonInput.From, jsonInput.To); !ans {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if existDoctor && existRoomInHospital { // если все прошло хорошо будет обновленое расписание
			userRepo := database.NewGormUserRepository()
			database.UpdateDataInTimetableByID(userRepo, idTimetable, jsonInput)
			c.AbortWithStatus(http.StatusOK) //mb change
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

func DeleteRecordFromTimetable() gin.HandlerFunc {
	return func(c *gin.Context) {
		//only for admin or manager
		roles := service.Authorization(c)
		if !strings.Contains(roles, "admin") && !strings.Contains(roles, "manager") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		idTimetable := c.Param("id")
		_, err := strconv.Atoi(idTimetable)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		userRepo := database.NewGormUserRepository()
		database.DeleteTimetableByID(userRepo, idTimetable)
		c.AbortWithStatus(http.StatusOK) //mb change
	}
}

func DeleteTimetableForDoctor() gin.HandlerFunc {
	return func(c *gin.Context) {
		//only for admin or manager
		roles := service.Authorization(c)
		if !strings.Contains(roles, "admin") && !strings.Contains(roles, "manager") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		idDoctor := c.Param("id")
		_, err := strconv.Atoi(idDoctor)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		userRepo := database.NewGormUserRepository()
		database.DeleteTimetableForDoctorByID(userRepo, idDoctor)
		c.AbortWithStatus(http.StatusOK) //mb change
	}
}

func DeleteTimetableForHospital() gin.HandlerFunc {
	return func(c *gin.Context) {
		//only for admin or manager
		roles := service.Authorization(c)
		if !strings.Contains(roles, "admin") && !strings.Contains(roles, "manager") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		idHospital := c.Param("id")
		_, err := strconv.Atoi(idHospital)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		userRepo := database.NewGormUserRepository()
		database.DeleteTimetableForHospitalByID(userRepo, idHospital)
		c.AbortWithStatus(http.StatusOK) //mb change
	}
}

func GetTimetableByIdHospital() gin.HandlerFunc {
	return func(c *gin.Context) {
		idHospital := c.Param("id")
		_, err := strconv.Atoi(idHospital)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		//get data form query
		from := c.Query("from")
		to := c.Query("to")
		layout := "2006-01-02T15:04:05Z07:00"
		fromTime, err := time.Parse(layout, from)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		toTime, err := time.Parse(layout, to)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		userRepo := database.NewGormUserRepository()
		timeTables := database.GetTimetableByIdHospitalInSegment(userRepo, idHospital, fromTime, toTime)
		c.JSON(http.StatusOK, timeTables)
	}
}

func GetTimetableByIdDoctor() gin.HandlerFunc {
	return func(c *gin.Context) {
		idDoctor := c.Param("id")
		_, err := strconv.Atoi(idDoctor)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		//get data form query
		from := c.Query("from")
		to := c.Query("to")
		layout := "2006-01-02T15:04:05Z07:00"
		fromTime, err := time.Parse(layout, from)
		if err != nil {
			fmt.Println("error 1", err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		toTime, err := time.Parse(layout, to)
		if err != nil {
			fmt.Println("error 2", err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		userRepo := database.NewGormUserRepository()
		timeTables := database.GetTimetableByIdDoctorInSegment(userRepo, idDoctor, fromTime, toTime)
		c.JSON(http.StatusOK, timeTables)
	}
}

func GetTimetableInHospitalRoom() gin.HandlerFunc {
	return func(c *gin.Context) {
		//only for admin or manager or doctor
		roles := service.Authorization(c)
		if !strings.Contains(roles, "admin") && !strings.Contains(roles, "manager") && !strings.Contains(roles, "doctor") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		idHospital := c.Param("id")
		if _, err := strconv.Atoi(idHospital); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		room := c.Param("room")

		//get data form query
		from := c.Query("from")
		to := c.Query("to")
		layout := "2006-01-02T15:04:05Z07:00"
		fromTime, err := time.Parse(layout, from)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		toTime, err := time.Parse(layout, to)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		userRepo := database.NewGormUserRepository()
		timeTables := database.GetTimetableByIdHospitalAndRoomForAMD(userRepo, idHospital, room, fromTime, toTime)
		c.JSON(http.StatusOK, timeTables)
	}
}
