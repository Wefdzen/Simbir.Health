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

// доступ к записям у которых pacientId eq or doctorId eq
func GetHistoryOfVisits() gin.HandlerFunc {
	return func(c *gin.Context) {
		idPacient := c.Param("id")
		idUserFromJwt := service.GetIdOfUserByJWT(c)
		idUser, err := strconv.Atoi(idUserFromJwt)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		roles := service.Authorization(c)
		pacientId, err := strconv.Atoi(idPacient)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if idUser != pacientId && !strings.Contains(roles, "user") {
			//проверка doctor
			if !strings.Contains(roles, "doctor") {
				c.AbortWithStatus(http.StatusUnauthorized)
				fmt.Println("here error")
				return
			}
		} //good
		userRepo := database.NewGormUserRepository()
		result := database.GetListHistoryByIdPacient(userRepo, idPacient)
		c.JSON(http.StatusOK, result)
	}
}

func GetHistory() gin.HandlerFunc {
	return func(c *gin.Context) {
		idHistory := c.Param("id")
		idUserFromJwt := service.GetIdOfUserByJWT(c)
		idUser, err := strconv.Atoi(idUserFromJwt)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		roles := service.Authorization(c)
		userRepo := database.NewGormUserRepository()
		result := database.GetListHistoryByIdHistory(userRepo, idHistory)
		if result.Room == "" {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		//проверка на пользователя
		if result.PacientId != idUser {
			//проверка doctor
			if !strings.Contains(roles, "doctor") {
				c.AbortWithStatus(http.StatusUnauthorized)
				fmt.Println("here error")
				return
			}
		}
		c.JSON(http.StatusOK, result)
	}
}

func CreateNewHistoryVisit() gin.HandlerFunc {
	return func(c *gin.Context) {
		var jsonInput database.History
		if err := c.BindJSON(&jsonInput); err != nil {
			fmt.Println("here 1")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		idUserFromJwt := service.GetIdOfUserByJWT(c)
		idUser, err := strconv.Atoi(idUserFromJwt)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		roles := service.Authorization(c)
		if idUser != jsonInput.PacientId && !strings.Contains(roles, "user") {
			//проверка на admin, manager, doctor
			if !strings.Contains(roles, "admin") && !strings.Contains(roles, "manager") && !strings.Contains(roles, "doctor") {
				c.AbortWithStatus(http.StatusUnauthorized)
				fmt.Println("here error")
				return
			}

		} //good
		//проверка сущностей
		doctorId := strconv.Itoa(jsonInput.DoctorId)
		if ok, _ := service.CheckExistDoctor(doctorId, c); !ok {
			c.AbortWithStatus(http.StatusBadRequest)
			fmt.Println("here 2")
			return
		}
		hospitalId := strconv.Itoa(jsonInput.HospitalId)
		fmt.Println(hospitalId)
		if ok, _ := service.CheckExistRoomInHospital(jsonInput.Room, hospitalId, c); !ok {
			c.AbortWithStatus(http.StatusBadRequest)
			fmt.Println("here 3")
			return
		}
		//create a new history
		userRepo := database.NewGormUserRepository()
		database.CreateHistoryOfVisit(userRepo, &jsonInput)
		c.JSON(http.StatusOK, gin.H{
			"status": "create success",
		})
	}
}

func UpdateNewHistoryVisit() gin.HandlerFunc {
	return func(c *gin.Context) {
		idHistory := c.Param("id")
		var jsonInput database.History
		if err := c.BindJSON(&jsonInput); err != nil {
			fmt.Println("here 1")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		idUserFromJwt := service.GetIdOfUserByJWT(c)
		idUser, err := strconv.Atoi(idUserFromJwt)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		roles := service.Authorization(c)
		if idUser != jsonInput.PacientId && !strings.Contains(roles, "user") {
			//проверка на admin, manager, doctor
			if !strings.Contains(roles, "admin") && !strings.Contains(roles, "manager") && !strings.Contains(roles, "doctor") {
				c.AbortWithStatus(http.StatusUnauthorized)
				fmt.Println("here error")
				return
			}

		} //good
		//проверка сущностей
		doctorId := strconv.Itoa(jsonInput.DoctorId)
		if ok, _ := service.CheckExistDoctor(doctorId, c); !ok {
			c.AbortWithStatus(http.StatusBadRequest)
			fmt.Println("here 2")
			return
		}
		hospitalId := strconv.Itoa(jsonInput.HospitalId)
		fmt.Println(hospitalId)
		if ok, _ := service.CheckExistRoomInHospital(jsonInput.Room, hospitalId, c); !ok {
			c.AbortWithStatus(http.StatusBadRequest)
			fmt.Println("here 3")
			return
		}
		//create a new history
		userRepo := database.NewGormUserRepository()
		database.UpdateHistoryOfVisit(userRepo, idHistory, &jsonInput)
		c.JSON(http.StatusOK, gin.H{
			"status": "update success",
		})
	}
}
