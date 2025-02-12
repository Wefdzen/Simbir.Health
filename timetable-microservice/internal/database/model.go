package database

import "time"

// delete will be a not soft
type Timetable struct {
	ID         uint
	HospitalId int       `json:"hospitalId"`
	DoctorId   int       `json:"doctorId"`
	From       time.Time `json:"from"` //как я понял просто промежуток времяни для записи к доктору.
	To         time.Time `json:"to"`
	Room       string    `json:"room"`
}

// talon
type Appointment struct {
	ID          uint
	TimetableId int       //uid of расписания
	ClientId    int       //чел который записался
	Time        time.Time // время записи после этого оно не доступно
}

type RequestAppointmentByTime struct {
	Time string `json:"time"`
}
