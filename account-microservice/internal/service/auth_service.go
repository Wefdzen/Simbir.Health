package service

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type TokensCouple struct {
	AccessToken  string
	RefreshToken string
}

func GenerateTokensCouple(uidUser string, origRoles []string) TokensCouple {
	//create JWT tokens
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        uidUser,                                 //id from bd
		"liveToken": time.Now().Add(time.Minute * 15).Unix(), // 15 min Я знаю про exp
		"roles":     origRoles,                               // по умол. user а уже ост с помощью админов
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("secret_key")))
	if err != nil {
		log.Fatal(err)
	}

	token2 := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        uidUser,                                    //id from bd
		"liveToken": time.Now().Add(time.Hour * 24 * 30).Unix(), // 30 day
	})
	tokenString2, err := token2.SignedString([]byte(os.Getenv("secret_key")))
	if err != nil {
		log.Fatal(err)
	}

	//couple of tokens
	tmp := TokensCouple{
		AccessToken:  tokenString,
		RefreshToken: tokenString2,
	}

	return tmp
}

func CheckRefreshToken(c *gin.Context, RefreshTokenOutUser, OrigRefreshToken string) bool {

	if RefreshTokenOutUser != OrigRefreshToken {
		return false
	}

	// Парсим refresh token
	token, err := jwt.Parse(RefreshTokenOutUser, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("secret_key")), nil
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return false
	}

	// Если refresh token истек, возвращаем false
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) >= claims["liveToken"].(float64) {
			return false
		}
	} else {
		return false
	}

	return true //refresh token success!
}

func GetIDFromToken(c *gin.Context, accessToken string) string {
	//проверять уже бесмысленно он до этого прошео middleware но получить claims надо
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("secret_key")), nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims["id"].(string)
	}
	return "-1"
}
