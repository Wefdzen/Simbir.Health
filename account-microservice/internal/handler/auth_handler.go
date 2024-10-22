package handler

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"wefdzen/internal/database"
	"wefdzen/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func SignUp() gin.HandlerFunc {
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
		database.RegisterUser(userRepo, &jsonInput, []string{"user"}) //по умолчани/ все будут user

		c.JSON(http.StatusOK, gin.H{
			"status": "registration a new user success",
		})
	}
}

func SignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var jsonInput database.User
		if err := c.BindJSON(&jsonInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		userRepo := database.NewGormUserRepository()
		if ok := database.CheckPassword(userRepo, &jsonInput); !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "incorrect password or username",
			})
			return // exit => don't pass jwt couple
		}
		//id of user from db for access jwt
		idUser := database.GetID(userRepo, &jsonInput)
		userId := fmt.Sprintf("%v", idUser) // Преобразование id в строку
		origRoles := database.GetRoles(userRepo, userId)
		tokens := service.GenerateTokensCouple(userId, origRoles)
		//set to database a refreshToken
		database.SetRefToken(userRepo, userId, tokens.RefreshToken)
		//Set in cookie
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("accessToken", tokens.AccessToken, 3600*24*30, "", "", false, true)
		c.SetCookie("refreshToken", tokens.RefreshToken, 3600*24*30, "", "", false, true)
		c.JSON(http.StatusOK, gin.H{
			"status": "login success",
		})
	}
}

func SignOut() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("accessToken", "", -1, "", "", false, true)  //сразу сдох
		c.SetCookie("refreshToken", "", -1, "", "", false, true) //сразу сдох
		c.JSON(http.StatusOK, gin.H{
			"status": "cookie delete", // tipo signout
		})
	}
}

type AccessRequest struct {
	AccessToken string `json:"accessToken"`
}

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("accessToken")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		//parsing token
		token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			// Проверяем метод подписи
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// Возвращаем секретный ключ
			return []byte(os.Getenv("secret_key")), nil
		})
		if err != nil { //какойто не такой вообще токен
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err,
			})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Проверяем вручную exp = liveToken поле Parse сам проверяет exp
			if float64(time.Now().Unix()) >= claims["liveToken"].(float64) {
				// Если claims не валидны
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			} else {
				//accessToken прокатил
				c.JSON(http.StatusOK, gin.H{"status": "validate success"})
				return
			}
		} else {
			// Если claims не валидны
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}

func Refresh() gin.HandlerFunc {
	return func(c *gin.Context) {
		var jsonInput RefreshRequest
		if err := c.BindJSON(&jsonInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//если чел не передал token беру из cookie
		var err error
		if jsonInput.RefreshToken == "" {
			jsonInput.RefreshToken, err = c.Cookie("refreshToken")
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}

		//parsing token
		token, err := jwt.Parse(jsonInput.RefreshToken, func(token *jwt.Token) (interface{}, error) {
			// Проверяем метод подписи
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// Возвращаем секретный ключ
			return []byte(os.Getenv("secret_key")), nil
		})
		if err != nil { //какойто не такой вообще токен
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err,
			})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Проверяем вручную exp = liveToken поле Parse сам проверяет exp
			if float64(time.Now().Unix()) >= claims["liveToken"].(float64) {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			} else {
				//get ref from db
				userRepo := database.NewGormUserRepository()
				userId := fmt.Sprintf("%v", claims["id"]) // Преобразование id в строку
				origRefTok := database.GetRefToken(userRepo, userId)
				//checkRefresh from db and his exp
				if !service.CheckRefreshToken(c, jsonInput.RefreshToken, origRefTok) {
					c.AbortWithStatus(http.StatusUnauthorized)
					return
				} else {
					//get role by id in db for regenerate tokens
					origRoles := database.GetRoles(userRepo, userId)
					tokens := service.GenerateTokensCouple(userId, origRoles)
					//set new ref token to db
					database.SetRefToken(userRepo, userId, tokens.RefreshToken)
					//set a new cookie
					c.SetSameSite(http.SameSiteLaxMode)
					c.SetCookie("accessToken", tokens.AccessToken, 3600*24*30, "", "", false, true)
					c.SetCookie("refreshToken", tokens.RefreshToken, 3600*24*30, "", "", false, true)
					c.JSON(http.StatusOK, gin.H{}) // all good
				}

			}
		} else {
			// Если claims не валидны
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
