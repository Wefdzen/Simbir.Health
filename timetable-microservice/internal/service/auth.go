package service

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authorization(c *gin.Context) string {
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
		r := fmt.Sprintf("%v", roles)
		return r
	}
	return ""
}
