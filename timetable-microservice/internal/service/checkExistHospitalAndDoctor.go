package service

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckExistDoctor(presumablyIdDoctor string, c *gin.Context) (bool, error) {
	accessToken, err := c.Cookie("accessToken")
	if err != nil {
		return false, err
	}
	// Создание HTTP-запроса к сервису аутентификации
	client := &http.Client{
		Timeout: 5 * time.Second, // Установим таймаут на запрос
	}
	// Формируем запрос к другому микросервису
	req, err := http.NewRequest("GET", "http://localhost:8080/api/Doctors/Exist", nil)
	if err != nil {
		return false, err
	}
	req.AddCookie(&http.Cookie{Name: "idDoctor", Value: presumablyIdDoctor})
	req.AddCookie(&http.Cookie{Name: "accessToken", Value: accessToken})

	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// Проверка кода статуса ответа
	if resp.StatusCode != http.StatusOK {
		return false, nil
	}
	return true, nil //doctor exist
}

func CheckExistRoomInHospital(room, presumablyIdHospital string, c *gin.Context) (bool, error) {
	accessToken, err := c.Cookie("accessToken")
	if err != nil {
		return false, err
	}
	// Создание HTTP-запроса к сервису аутентификации
	client := &http.Client{
		Timeout: 5 * time.Second, // Установим таймаут на запрос
	}
	// Формируем запрос к другому микросервису
	req, err := http.NewRequest("GET", "http://localhost:8081/api/Hospitals/Exist", nil)
	if err != nil {
		return false, err
	}
	//Короче room токо на английских и цифрах
	req.AddCookie(&http.Cookie{Name: "room", Value: room})
	req.AddCookie(&http.Cookie{Name: "idHospital", Value: presumablyIdHospital})
	req.AddCookie(&http.Cookie{Name: "accessToken", Value: accessToken})
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// Проверка кода статуса ответа
	if resp.StatusCode != http.StatusOK {
		return false, nil
	}
	return true, nil //doctor exist
}
