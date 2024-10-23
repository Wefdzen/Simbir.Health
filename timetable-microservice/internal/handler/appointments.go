package handler

import (
	"net/http"
	"strconv"
	"strings"
	"time"
	"wefdzen/internal/database"
	"wefdzen/internal/service"

	"github.com/gin-gonic/gin"
)

func GetFreeAppointments() gin.HandlerFunc {
	return func(c *gin.Context) {
		idTimetable := c.Param("id")
		if _, err := strconv.Atoi(idTimetable); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		// Получить расписание по ID
		userRepo := database.NewGormUserRepository()
		timetable := database.GetFreeAppointmentsById(userRepo, idTimetable)
		if timetable.HospitalId == 0 { // если расписания не существует
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		// Генерируем все возможные талоны по 30 минут
		timetable.From = timetable.From.UTC() // +3msk -
		timetable.To = timetable.To.UTC()
		freeSlots := service.GenerateAppointmentsSlots(timetable.From, timetable.To, userRepo)

		c.JSON(http.StatusOK, freeSlots)
	}
}

func RecordingToAppointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		idTimetable := c.Param("id")
		if _, err := strconv.Atoi(idTimetable); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		var jsonInput database.RequestAppointmentByTime
		if err := c.BindJSON(&jsonInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		layout := "2006-01-02T15:04:05Z07:00"
		time, err := time.Parse(layout, jsonInput.Time)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		//Checking paramaters time in body from req 30 min
		service.CheckMultiplyOfTimeOne(time, c)

		idOfTimetable, err := strconv.Atoi(idTimetable)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		idOfClientFromJWT := service.GetIdOfUserByJWT(c) //может вернуть пустую строку просто упадет на next step
		idClient, err := strconv.Atoi(idOfClientFromJWT)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		appointmentForCreate := database.Appointment{
			TimetableId: idOfTimetable,
			ClientId:    idClient,
			Time:        time,
		}

		userRepo := database.NewGormUserRepository()
		statusOfTime := database.CheckAvailibleOfTimeInTimetableByTime(userRepo, time)
		if !statusOfTime {
			// если время занято
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		//create a new appointment
		database.CreateNewAppointmentInTimetableAll(userRepo, &appointmentForCreate)
		c.AbortWithStatus(http.StatusOK)
	}
}

func CancelingAppointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		idAppointment := c.Param("id")
		if _, err := strconv.Atoi(idAppointment); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		//for admin, manager and how create a appointment
		idClient := service.GetIdOfUserByJWT(c)

		//проверка usera его ли это id
		userRepo := database.NewGormUserRepository()
		if ok := database.CheckThisClientCreateThisAppointment(userRepo, idClient, idAppointment); !ok {
			//если с user не прокатило то чекаем уже роли
			roles := service.Authorization(c)
			if !strings.Contains(roles, "admin") && !strings.Contains(roles, "manager") {
				c.AbortWithStatus(http.StatusForbidden)
				return
			}
		} //all good access success to delete
		database.DeleteDataAppointmentById(userRepo, idAppointment)
		c.JSON(http.StatusOK, gin.H{
			"status": "delete success",
		})
	}
}
