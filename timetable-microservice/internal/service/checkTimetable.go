package service

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckMultiplyOfTime(timeFrom, timeTo time.Time, c *gin.Context) bool {
	// Проверяем, кратно ли время 30 минутам для обоих значений (timeFrom и timeTo)
	if timeFrom.Minute()%30 != 0 || timeTo.Minute()%30 != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "the time should be a multiple of 30 minutes"})
		return false
	}
	// Проверяем, что секунды равны 0 у обоих значений
	if timeFrom.Second() != 0 || timeTo.Second() != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "the second should be equal to 0"})
		return false
	}
	return true
}

func CheckMultiplyOfTimeOne(time time.Time, c *gin.Context) {
	// Проверяем, кратно ли время 30 минутам для обоих значений (timeFrom и timeTo)
	if time.Minute()%30 != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "the time should be a multiple of 30 minutes"})
		return
	}
	// Проверяем, что секунды равны 0 у обоих значений
	if time.Second() != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "the second should be equal to 0"})
		return
	}
}

func CheckTimeDifference(timeFrom, timeTo time.Time, c *gin.Context) bool {
	duration := timeTo.Sub(timeFrom)
	// Проверяем, что время "to" больше "from"
	if duration <= 0 {
		return false
	}
	// Проверяем, что разница не превышает 12 часов
	if duration.Hours() > 12 {
		return false
	}
	return true
}
