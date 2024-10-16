package database

import "time"

type UserRepository interface {
	CreateNewTimetableForDoctor(newTimetable *Timetable)
	UpdateDataInTimetable(idTimetable string, newTimetable Timetable)
	DeleteDataTimetable(idTimetable string)
	DeleteDataTimetableForDoctor(idDoctor string)
	DeleteDataTimetableForHospital(idHospital string)
	GetTimetableByIdHospital(idHospital string, from, to time.Time) []Timetable
	GetTimetableByIdDoctor(idDoctor string, from, to time.Time) []Timetable
	GetTimetableByIdHospitalAndRoom(idHospital, room string, form, to time.Time) []Timetable

	CreateNewAppointmentInTimetable(newAppointment *Appointment)
	CheckAvailibleOfTimeInTimetable(time time.Time) bool
	CheckThisClientCreateThisAppo(idClient, idAppointment string) bool
	DeleteDataAppointment(idAppointment string)
	GetFreeAppointments(idTimetable string) Timetable
}

func CreateTimetableForDoctor(repo UserRepository, newTimetable *Timetable) {
	repo.CreateNewTimetableForDoctor(newTimetable)
}

func UpdateDataInTimetableByID(repo UserRepository, idTimetable string, newTimetable Timetable) {
	repo.UpdateDataInTimetable(idTimetable, newTimetable)
}

func DeleteTimetableByID(repo UserRepository, idTimetable string) {
	repo.DeleteDataTimetable(idTimetable)
}

func DeleteTimetableForDoctorByID(repo UserRepository, idDoctor string) {
	repo.DeleteDataTimetableForDoctor(idDoctor)
}

func DeleteTimetableForHospitalByID(repo UserRepository, idHospital string) {
	repo.DeleteDataTimetableForHospital(idHospital)
}

func GetTimetableByIdHospitalInSegment(repo UserRepository, idHospital string, from, to time.Time) []Timetable {
	return repo.GetTimetableByIdHospital(idHospital, from, to)
}

func GetTimetableByIdDoctorInSegment(repo UserRepository, idDoctor string, from, to time.Time) []Timetable {
	return repo.GetTimetableByIdDoctor(idDoctor, from, to)
}

// A - admin, M - manager, D - doctor.
func GetTimetableByIdHospitalAndRoomForAMD(repo UserRepository, idHospital, room string, from, to time.Time) []Timetable {
	return repo.GetTimetableByIdHospitalAndRoom(idHospital, room, from, to)
}
func CreateNewAppointmentInTimetableAll(repo UserRepository, newAppointment *Appointment) {
	repo.CreateNewAppointmentInTimetable(newAppointment)
}
func CheckAvailibleOfTimeInTimetableByTime(repo UserRepository, time time.Time) bool {
	return repo.CheckAvailibleOfTimeInTimetable(time)
}
func CheckThisClientCreateThisAppointment(repo UserRepository, idClient, idAppointment string) bool {
	return repo.CheckThisClientCreateThisAppo(idClient, idAppointment)
}
func DeleteDataAppointmentById(repo UserRepository, idAppointment string) {
	repo.DeleteDataAppointment(idAppointment)
}
func GetFreeAppointmentsById(repo UserRepository, idTimetable string) Timetable {
	return repo.GetFreeAppointments(idTimetable)
}
