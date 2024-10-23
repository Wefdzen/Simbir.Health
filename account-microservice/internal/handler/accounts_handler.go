package handler

import (
	"net/http"
	"strconv"
	"strings"
	"wefdzen/internal/database"
	"wefdzen/internal/service"

	"github.com/gin-gonic/gin"
)

func Me() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("accessToken")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		idUser := service.GetIDFromToken(c, accessToken)
		//end func
		userRepo := database.NewGormUserRepository()
		me := database.GetAllInfoByID(userRepo, idUser)
		c.JSON(http.StatusOK, me)
	}
}

func UpdateDataByUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("accessToken")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		idUser := service.GetIDFromToken(c, accessToken)
		userRepo := database.NewGormUserRepository()
		var jsonInput database.User
		if err := c.BindJSON(&jsonInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		database.UpdateDataAccount(userRepo, idUser, jsonInput)
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}

func GetAllAccounts() gin.HandlerFunc {
	return func(c *gin.Context) {
		//only for admin
		roles := service.Authorization(c)
		if !strings.Contains(roles, "admin") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		//get data form query
		from := c.Query("from")
		if from == "" {
			from = "0" // будет по умолчанию с самого начала
		}
		count := c.Query("count")
		if count == "" {
			count = "10" // 10 записей
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
		allUsers := database.GetAllInfoAllAccounts(userRepo, fromI, countI)
		c.JSON(http.StatusOK, allUsers)
	}
}

func CreateAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		//only for admin
		roles := service.Authorization(c)
		if !strings.Contains(roles, "admin") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		var jsonInput database.User
		if err := c.BindJSON(&jsonInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		//hashing password
		tmpPassword := jsonInput.Password
		jsonInput.Password, _ = service.HashPassword(tmpPassword)

		//connect to db
		userRepo := database.NewGormUserRepository()
		database.NewAccountByAdmin(userRepo, &jsonInput)
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}

func UpdateDataByAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		//only for admin
		roles := service.Authorization(c)
		if !strings.Contains(roles, "admin") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		idUser := c.Param("id")
		userRepo := database.NewGormUserRepository()
		var jsonInput database.User
		if err := c.BindJSON(&jsonInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		database.UpdateDataAccountAdmin(userRepo, idUser, jsonInput)
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}

func DeleteAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		//only for admin
		roles := service.Authorization(c)
		if !strings.Contains(roles, "admin") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		//Get ID for soft delete
		idUser := c.Param("id")
		userRepo := database.NewGormUserRepository()

		//если ошибка по айди не нашел то NotFound
		err := database.SoftDeleteAccountAdmin(userRepo, idUser)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		//delete success
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}
