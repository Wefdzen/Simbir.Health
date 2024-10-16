package service

import (
	"time"
	"wefdzen/internal/database"
)

func GenerateAppointmentsSlots(from time.Time, to time.Time, repo *database.GormUserRepository) []time.Time {
	var freeSlots []time.Time
	currentTime := from

	// Генерация слотов каждые 30 минут, пока время не превысит конечное
	for currentTime.Before(to) || currentTime.Equal(to) {
		if repo.CheckAvailibleOfTimeInTimetable(currentTime) {
			freeSlots = append(freeSlots, currentTime)
		}
		currentTime = currentTime.Add(30 * time.Minute)
	}

	return freeSlots
}
