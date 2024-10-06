package database

import "time"

type UserRepository interface {
	CreateNewTimetableForDoctor(newTimetable *Timetable)
	UpdateDataInTimetable(idTimetable string, newTimetable Timetable)
	DeleteDataTimetable(idTimetable string)
	DeleteDataTimetableForDoctor(idDoctor string)
	DeleteDataTimetableForHospital(idHospital string)
	GetTimetableByIdHospital(idHospital string, from, to time.Time) []Timetable
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
