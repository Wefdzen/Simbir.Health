package database

import "time"

type History struct {
	ID         uint
	Date       time.Time `json:"date"`
	PacientId  int       `json:"pacientId"`
	HospitalId int       `json:"hospitalId"`
	DoctorId   int       `json:"doctorId"`
	Room       string    `json:"room"`
	Data       string    `json:"data"`
}
