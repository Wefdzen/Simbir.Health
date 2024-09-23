package handler

import (
	"net/http"
	"strconv"
	"wefdzen/internal/database"
	"wefdzen/internal/service"

	"github.com/gin-gonic/gin"
)

func Me() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO create func get idUser in service
		accessToken, err := c.Cookie("accessToken")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		idUser := service.GetIDFromToken(c, accessToken)
		//end func
		userRepo := database.NewGormUserRepository()
		me := database.GetAllInfoByID(userRepo, idUser)
		c.JSON(http.StatusOK, gin.H{
			"lastName":  me.LastName,
			"firstName": me.FirstName,
			"username":  me.UserName,
			"roles":     me.Roles,
		})
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
		//get data form query
		from := c.Query("from")
		count := c.Query("count")
		fromI, _ := strconv.Atoi(from)
		countI, _ := strconv.Atoi(count)

		userRepo := database.NewGormUserRepository()
		allUsers := database.GetAllInfoAllAccounts(userRepo, fromI, countI)
		c.JSON(http.StatusOK, allUsers)
	}
}

func CreateAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
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
	}
}

func DeleteAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
