package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("accessToken")
		if err != nil {
			fmt.Println("1")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// Создание HTTP-запроса к сервису аутентификации
		client := &http.Client{
			Timeout: 5 * time.Second, // Установим таймаут на запрос
		}
		// Формируем запрос к другому микросервису
		req, err := http.NewRequest("GET", "http://localhost:8080/api/Authentication/Validate", nil)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		req.AddCookie(&http.Cookie{Name: "accessToken", Value: accessToken})

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Ошибка запроса к Authentication сервису:", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		defer resp.Body.Close()

		// Проверка кода статуса ответа
		if resp.StatusCode != http.StatusOK {
			fmt.Println("Authentication не прошла, статус:", resp.StatusCode)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}

// will be after middleware Authe
func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		// незнаю либо просто брать из accessToken либо брать(вроде сверка с бд ролей не надо)
		c.Next()
	}
}
