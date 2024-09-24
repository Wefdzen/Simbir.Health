package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

// will be after middleware Authentication
func Authorization(c *gin.Context) {
	accessToken, _ := c.Cookie("accessToken")

	//parsing token // err не надо до этого уже проверенно
	token, _ := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		// Проверяем метод подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Возвращаем секретный ключ
		return []byte(os.Getenv("secret_key")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		roles := claims["roles"]
		fmt.Println(roles)
		r := fmt.Sprintf("%v", roles)
		fmt.Println("R:", r)
	}
	c.Next()
}
